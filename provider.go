package givv

type Provider[T any] interface {
	Get() T
}
