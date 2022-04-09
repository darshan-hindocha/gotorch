package nn

import (
	"github.com/darshan-hindocha/gotorch/internal/tensor"
)


type LinearLayer struct {
	weight       tensor.Tensor
	bias         tensor.Tensor
	shape        []int
	in_features  int
	out_features int
}

func (ll *LinearLayer) Initialise(x, y LayerData) error {

	return nil
}

func (ll *LinearLayer) Reset_parameters() error {
	err := ll.weight.Initialise()
	if err != nil {
		return err
	}

	err = ll.bias.Initialise()
	if err != nil {
		return err
	}
	return nil
}

func (ll *LinearLayer) Forward() error {
	
	return nil
}
