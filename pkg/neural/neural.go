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

func Normalize(sensors *types.Sensors) [5]float32 {
	data := [5]float32{
		float32(sensors.Timestamp),
		sensors.Temperature,
		sensors.Pressure,
		sensors.Humidity,
		sensors.Gas,
	}
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

func Predict(sensors *types.Sensors) (*types.OptimizationData, error) {
	model, err := tf.LoadSavedModel(modelDir, []string{modelTag}, nil)
	if err != nil {
		return nil, err
	}
	defer model.Session.Close()

	data := Normalize(sensors)

	tensor, err := tf.NewTensor([][5]float32{data})
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
		Timestamp:        sensors.Timestamp,
		WindowsAreOpened: output[0] > treshold,
		PeopleInTheRoom:  output[1] > treshold,
	}

	return opt, nil
}
