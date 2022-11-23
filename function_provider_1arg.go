package givv

type Function1ArgProvider1Arg[T any, A any] struct {
	resolver *Resolver
	function func(A) T
}

func NewFunction1ArgProvider1Arg[T any, A any](resolver *Resolver, function func(A) T) Provider1Arg[T, A]{
	return Function1ArgProvider1Arg[T, A]{
		resolver, 
		function,
	}
}

func(provider Function1ArgProvider1Arg[T, A]) Get(arg A) T {
	return provider.function(arg)
}

// --------------

type Function2ArgsProvider1Arg[T any, FA any, PA any] struct {
	resolver *Resolver
	function func(FA, PA) T
	arg1 ArgSpec[FA]
	arg2 ArgSpec[PA]
}

func NewFunction2ArgsProvider1Arg[T any, FA any, PA any](
	resolver *Resolver, 
	function func(FA, PA) T, 
	arg1 ArgSpec[FA],
	arg2 ArgSpec[PA],
	) Provider1Arg[T, PA]{

	return Function2ArgsProvider1Arg[T, FA, PA] {
		resolver: resolver,
		function: function,
		arg1: arg1,
		arg2: arg2,
	}
}

func(provider Function2ArgsProvider1Arg[T, FA, PA]) Get(arg PA) T {
	val := provider.arg1.resolve(provider.resolver, newKeySlice())

	return provider.function(val, arg)
}

// --------------

type Function3ArgsProvider1Arg[T any, FA1 any, FA2 any, PA any] struct {
	resolver *Resolver
	function func(FA1, FA2, PA) T
	arg1 ArgSpec[FA1]
	arg2 ArgSpec[FA2]
	arg3 ArgSpec[PA]
}

func NewFunction3ArgsProvider1Arg[T any, FA1 any, FA2 any, PA any](
	resolver *Resolver, 
	function func(FA1, FA2, PA) T, 
	arg1 ArgSpec[FA1],
	arg2 ArgSpec[FA2],
	arg3 ArgSpec[PA],
	) Provider1Arg[T, PA]{

	return Function3ArgsProvider1Arg[T, FA1, FA2, PA] {
		resolver: resolver,
		function: function,
		arg1: arg1,
		arg2: arg2,
		arg3: arg3,
	}
}

func(provider Function3ArgsProvider1Arg[T, FA1, FA2, PA]) Get(arg PA) T {
	val1 := provider.arg1.resolve(provider.resolver, newKeySlice())
	val2 := provider.arg2.resolve(provider.resolver, newKeySlice())

	return provider.function(val1, val2, arg)
}

// --------------

type Function4ArgsProvider1Arg[T any, FA1 any, FA2 any, FA3 any, PA any] struct {
	resolver *Resolver
	function func(FA1, FA2, FA3, PA) T
	arg1 ArgSpec[FA1]
	arg2 ArgSpec[FA2]
	arg3 ArgSpec[FA3]
	arg4 ArgSpec[PA]
}

func NewFunction4ArgsProvider1Arg[T any, FA1 any, FA2 any, FA3 any, PA any](
	resolver *Resolver, 
	function func(FA1, FA2, FA3, PA) T, 
	arg1 ArgSpec[FA1],
	arg2 ArgSpec[FA2],
	arg3 ArgSpec[FA3],
	arg4 ArgSpec[PA],
	) Provider1Arg[T, PA]{

	return Function4ArgsProvider1Arg[T, FA1, FA2, FA3, PA] {
		resolver: resolver,
		function: function,
		arg1: arg1,
		arg2: arg2,
		arg3: arg3,
		arg4: arg4,
	}
}

func(provider Function4ArgsProvider1Arg[T, FA1, FA2, FA3, PA]) Get(arg PA) T {
	val1 := provider.arg1.resolve(provider.resolver, newKeySlice())
	val2 := provider.arg2.resolve(provider.resolver, newKeySlice())
	val3 := provider.arg3.resolve(provider.resolver, newKeySlice())

	return provider.function(val1, val2, val3, arg)
}

