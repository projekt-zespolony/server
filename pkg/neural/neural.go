package neural

import (
	"math"

	"github.com/projekt-zespolony/server/pkg/types"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

const (
	modelDir    = "model"
	modelTag    = "tag"
	treshold    = 0.5
	inputLayer  = "input_input"
	outputLayer = "output/Sigmoid"
)

func Normalize(data []float32) []float32 {
	dataLength := float32(len(data))

	sum := float32(0)
	for _, d := range data {
		sum += d
	}
	mean := sum / dataLength

	sumSquare := float32(0)
	for _, d := range data {
		sumSquare += (d - mean) * (d - mean)
	}
	variance := sumSquare / (dataLength - 1)

	stdev := float32(math.Sqrt(float64(variance)))

	for i := range data {
		data[i] = (data[i] - mean) / stdev
	}

	return data
}

func Predict(sensors []*types.Sensors) (*types.OptimizationData, error) {
	model, err := tf.LoadSavedModel(modelDir, []string{modelTag}, nil)
	if err != nil {
		return nil, err
	}
	defer model.Session.Close()

	timestamps := []float32{}
	temperatures := []float32{}
	pressures := []float32{}
	humidities := []float32{}
	gases := []float32{}
	for i := range sensors {
		timestamps = append(temperatures, float32(sensors[i].Timestamp))
		temperatures = append(temperatures, sensors[i].Temperature)
		pressures = append(pressures, sensors[i].Pressure)
		humidities = append(humidities, sensors[i].Humidity)
		gases = append(gases, sensors[i].Gas)
	}

	timestamps = Normalize(timestamps)
	temperatures = Normalize(temperatures)
	pressures = Normalize(pressures)
	humidities = Normalize(humidities)
	gases = Normalize(gases)

	allData := [][5]float32{}
	for i := range sensors {
		allData = append(allData, [5]float32{
			timestamps[i],
			temperatures[i],
			pressures[i],
			humidities[i],
			gases[i],
		})
	}

	tensor, err := tf.NewTensor(allData)
	if err != nil {
		return nil, err
	}

	result, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			model.Graph.Operation(inputLayer).Output(0): tensor,
		},
		[]tf.Output{
			model.Graph.Operation(outputLayer).Output(0),
		},
		nil,
	)
	if err != nil {
		return nil, err
	}

	output := result[0].Value().([][]float32)[0]

	opt := &types.OptimizationData{
		Timestamp:        sensors[len(sensors)-1].Timestamp,
		WindowsAreOpened: output[0] > treshold,
		PeopleInTheRoom:  output[1] > treshold,
	}

	return opt, nil
}
