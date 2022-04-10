package nn

import (
	"github.com/darshan-hindocha/gotorch/internal/tensor"
)


type LinearLayer struct {
	Weight       tensor.Tensor
	Bias         tensor.Tensor
	In_features  int
	Out_features int
}

func (l *LinearLayer) Initialise(x, y int) error {
	l.In_features = x
	l.Out_features = y

	l.Weight = tensor.Tensor{
		Shape: []int{l.In_features,l.Out_features}
		Values: []float64.
	}

	return nil
}

func (l *LinearLayer) Reset_parameters() error {
	err := l.weight.Initialise()
	if err != nil {
		return err
	}

	err = l.bias.Initialise()
	if err != nil {
		return err
	}
	return nil
}

func (l *LinearLayer) Forward() error {
	
	return nil
}
