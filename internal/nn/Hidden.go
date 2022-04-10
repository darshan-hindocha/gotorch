package nn

type HiddenLayer interface {
	Initialise(x, y int) error
	Reset_parameters() error
	Forward() error
}

type ActivationFunction interface {
	Forward() error
}
