package givv

type Function2ArgsProvider2Args[T any, A1 any, A2 any] struct {
	resolver *Resolver
	function func(A1, A2) T
}

func NewFunction2ArgsProvider2Args[T any, A1 any, A2 any](
	resolver *Resolver, 
	function func(A1, A2) T,
	) Provider2Args[T, A1, A2]{

	return Function2ArgsProvider2Args[T, A1, A2] {
		resolver: resolver,
		function: function,
	}
}

func(provider Function2ArgsProvider2Args[T, A1, A2]) Get(arg1 A1, arg2 A2) T {
	return provider.function(arg1, arg2)
}
