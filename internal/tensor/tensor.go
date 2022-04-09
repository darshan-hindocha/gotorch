package tensor

type TensorOperations interface {
	initialise()
	zeros(*Tensor, []int)
	reshape(*Tensor, []int)
	view(*Tensor, []int)
	squeeze(*Tensor, []int)
}

type Tensor struct {
	Values []float64
	Shape   []int
}


func (t *Tensor) Initialise() error {
	
	return nil
}
