package givv

type FunctionProvider[T any] struct {
	function func() T
}

func NewFunctionProvider[T any](function func() T) Provider[T] {
	return FunctionProvider[T]{
		function: function,
	}
}

func (provider FunctionProvider[T]) Get() T {
	return provider.function()
}

// ----------------

type Function1ArgProvider[T any, A any] struct {
	resolver *Resolver
	function func(A) T
	argSpec ArgSpec[A]
}

func NewFunction1ArgProvider[T any, A any](resolver *Resolver, function func(A) T, arg ArgSpec[A]) Provider[T] {
	return Function1ArgProvider[T, A]{
		resolver: resolver,
		function: function,
		argSpec: arg,
	}
}

func(provider Function1ArgProvider[T, A]) Get() T {
	val := provider.argSpec.resolve(provider.resolver, newKeySlice())
	return provider.function(val)
}

func (provider Function1ArgProvider[T, A]) getWithPreviousKeys(previousKeys []any) T {
	val := provider.argSpec.resolve(provider.resolver, previousKeys)
	return provider.function(val)	
}

// -----------------

type Function2ArgsProvider[T any, A1 any, A2 any] struct {
	resolver *Resolver
	function func(A1, A2) T
	arg1 ArgSpec[A1]
	arg2 ArgSpec[A2]
}

func NewFunction2ArgsProvider[T any, A1 any, A2 any](
	resolver *Resolver, 
	function func(A1, A2) T, 
	arg1 ArgSpec[A1],
	arg2 ArgSpec[A2],
	) Provider[T] {

	return Function2ArgsProvider[T, A1, A2]{
		resolver: resolver,
		function: function,
		arg1: arg1,
		arg2: arg2,
	}
}

func(provider Function2ArgsProvider[T, A1, A2]) Get() T {

	val1 := provider.arg1.resolve(provider.resolver, newKeySlice())
	val2 := provider.arg2.resolve(provider.resolver, newKeySlice())

	return provider.function(val1, val2)
}

func(provider Function2ArgsProvider[T, A1, A2]) getWithPreviousKeys(previousKeys []any) T {

	val1 := provider.arg1.resolve(provider.resolver, previousKeys)
	val2 := provider.arg2.resolve(provider.resolver, previousKeys)

	return provider.function(val1, val2)
}

// -----------------

type Function3ArgsProvider[T any, A1 any, A2 any, A3 any] struct {
	resolver *Resolver
	function func(A1, A2, A3) T
	arg1 ArgSpec[A1]
	arg2 ArgSpec[A2]
	arg3 ArgSpec[A3]
}

func NewFunction3ArgsProvider[T any, A1 any, A2 any, A3 any](
	resolver *Resolver, 
	function func(A1, A2, A3) T, 
	arg1 ArgSpec[A1],
	arg2 ArgSpec[A2],
	arg3 ArgSpec[A3],
	) Provider[T] {

	return Function3ArgsProvider[T, A1, A2, A3]{
		resolver: resolver,
		function: function,
		arg1: arg1,
		arg2: arg2,
		arg3: arg3,
	}
}

func(provider Function3ArgsProvider[T, A1, A2, A3]) Get() T {

	val1 := provider.arg1.resolve(provider.resolver, newKeySlice())
	val2 := provider.arg2.resolve(provider.resolver, newKeySlice())
	val3 := provider.arg3.resolve(provider.resolver, newKeySlice())

	return provider.function(val1, val2, val3)
}

func(provider Function3ArgsProvider[T, A1, A2, A3]) getWithPreviousKeys(previousKeys []any) T {

	val1 := provider.arg1.resolve(provider.resolver, previousKeys)
	val2 := provider.arg2.resolve(provider.resolver, previousKeys)
	val3 := provider.arg3.resolve(provider.resolver, previousKeys)

	return provider.function(val1, val2, val3)
}

// -----------------

type Function4ArgsProvider[T any, A1 any, A2 any, A3 any, A4 any] struct {
	resolver *Resolver
	function func(A1, A2, A3, A4) T
	arg1 ArgSpec[A1]
	arg2 ArgSpec[A2]
	arg3 ArgSpec[A3]
	arg4 ArgSpec[A4]
}

func NewFunction4ArgsProvider[T any, A1 any, A2 any, A3 any, A4 any](
	resolver *Resolver, 
	function func(A1, A2, A3, A4) T, 
	arg1 ArgSpec[A1],
	arg2 ArgSpec[A2],
	arg3 ArgSpec[A3],
	arg4 ArgSpec[A4],
	) Provider[T] {

	return Function4ArgsProvider[T, A1, A2, A3, A4]{
		resolver: resolver,
		function: function,
		arg1: arg1,
		arg2: arg2,
		arg3: arg3,
		arg4: arg4,
	}
}

func(provider Function4ArgsProvider[T, A1, A2, A3, A4]) Get() T {

	val1 := provider.arg1.resolve(provider.resolver, newKeySlice())
	val2 := provider.arg2.resolve(provider.resolver, newKeySlice())
	val3 := provider.arg3.resolve(provider.resolver, newKeySlice())
	val4 := provider.arg4.resolve(provider.resolver, newKeySlice())

	return provider.function(val1, val2, val3, val4)
}

func(provider Function4ArgsProvider[T, A1, A2, A3, A4]) getWithPreviousKeys(previousKeys []any) T {

	val1 := provider.arg1.resolve(provider.resolver, previousKeys)
	val2 := provider.arg2.resolve(provider.resolver, previousKeys)
	val3 := provider.arg3.resolve(provider.resolver, previousKeys)
	val4 := provider.arg4.resolve(provider.resolver, previousKeys)

	return provider.function(val1, val2, val3, val4)
}









