package main

import (
	"github.com/darshan-hindocha/gotorch/internal/nn"
	"github.com/darshan-hindocha/gotorch/internal/tensor"
	"github.com/sirupsen/logrus"
)

func main() {

	model := Model{}

	err := model.InitialiseModel()
	if err != nil {
		logrus.Errorf("model initialisation error: %s", err)
	}
	logrus.Info("Main Ran!")
}

type Model struct {
	InputTensor  tensor.Tensor
	OutputTensor tensor.Tensor
	Layers       map[int]nn.HiddenLayer
	Activations  map[int]nn.ActivationFunction
	Shapes       map[int]int
}

func (model *Model) InitialiseModel() error {

	shapes := map[int]int{
		0: 32,
		1: 10,
		2: 10,
		3: 9,
	}

	layers := map[int]nn.HiddenLayer{
		1: new(nn.LinearLayer),
		2: new(nn.LinearLayer),
	}

	for index, layer := range layers {
		err := layer.Initialise(shapes[index-1], shapes[index])
		if err != nil {
			logrus.Errorf("layer initialise error: %s", err)
		}
	}
	activations := map[int]nn.ActivationFunction{
		1: new(nn.ReLu),
		2: new(nn.ReLu),
	}

	model.Layers = layers
	model.Activations = activations

	return nil
}
