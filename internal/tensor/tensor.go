package tensor

type TensorOperations interface {
	initialise()
	Zeros()
	Shape()
	reshape(*Tensor, []int)
	view(*Tensor, []int)
	squeeze(*Tensor, []int)
}

type Tensor struct {
	Values     map[int][]float32
	Dimensions int
	Shape      map[int]int
}

func (t *Tensor) Initialise() error {

	return nil
}


func (t *Tensor) Zeros(tensorShape map[int]int) error {
	t.Shape = tensorShape
	t.Dimensions = len(tensorShape)

	values := make(map[int][]float32)

	for i:=0; i < t.Dimensions; i++ {
		array := [10]float32
		var arraySlice []float32 = array
		values[i] = arraySlice
	}	

}


func (t *Tensor) Shape() error {

	return t.Shape
}