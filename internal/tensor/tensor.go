package tensor

type TensorInterface interface {
	zeros(*TensorStruct, []int)
	reshape(*TensorStruct, []int)
}

type TensorStruct struct {
	Tensor []float64
	Shape  []int
}
