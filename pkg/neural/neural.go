package neural

import (
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

func Predict(sensors *types.Sensors) (*types.OptimizationData, error) {
	model, err := tf.LoadSavedModel(modelDir, []string{modelTag}, nil)
	if err != nil {
		return nil, err
	}
	defer model.Session.Close()

	data := [5]float32{
		0,
		sensors.Temperature,
		sensors.Pressure,
		sensors.Humidity,
		sensors.Gas,
	}

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
