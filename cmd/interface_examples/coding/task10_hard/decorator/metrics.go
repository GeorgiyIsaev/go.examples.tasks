package decorator

type Metrics interface {
	GetCalls() int
	SetCalls() int
	GetErrors() int
	SetErrors() int
}
