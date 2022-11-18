package givv

type Provider[T any] interface {
	Get() T
}

type Provider1Arg[T any, A any] interface {
	Get(A) T
}

type Provider2Args[T any, A1 any, A2 any] interface {
	Get(A1, A2) T
}

type Provider3Args[T any, A1 any, A2 any, A3 any] interface {
	Get(A1, A2, A3) T
}

type Provider4Args[T any, A1 any, A2 any, A3 any, A4 any] interface {
	Get(A1, A2, A3, A4) T
}

type Provider5Args[T any, A1 any, A2 any, A3 any, A4 any, A5 any] interface {
	Get(A1, A2, A3, A4, A5) T
}

