package nn


type LinearLayer struct {
	weight       []float64
	bias         []float64
	shape        []int
	in_features  int
	out_features int
}

func (ll *LinearLayer) Reset_parameters() error {
	
	return nil
}

func (ll *LinearLayer) Forward() error {
	
	return nil
}
