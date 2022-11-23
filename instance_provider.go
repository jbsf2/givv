package givv

type instanceProvider[T any] struct {
	instance T
}

func newInstanceProvider[T any](instance T) Provider[T] {
	return instanceProvider[T]{
		instance: instance,
	}
}

func (provider instanceProvider[T]) Get() T {
	return provider.instance
}
