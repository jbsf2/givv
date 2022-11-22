package givv

type InstanceProvider[T any] struct {
	instance T
}

func NewInstanceProvider[T any](instance T) Provider[T] {
	return InstanceProvider[T]{
		instance: instance,
	}
}

func (instanceProvider InstanceProvider[T]) Get() T {
	return instanceProvider.instance
}
