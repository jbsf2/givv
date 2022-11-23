package givv

import (
	"reflect"
)

type providerWithCycleDetection[T any] interface {
	getWithPreviousKeys(previousKeys []any) T
}

type Resolver struct {
	providers map[any]any
}

func NewResolver() *Resolver {
	return &Resolver{
		providers: map[any]any{},
	}
}

func Resolve[T any, K any](resolver *Resolver, key key[T, K]) T {
	return resolveWithCycleDetection(resolver, key, []any{})
}

func resolveWithCycleDetection[T any, K any](resolver *Resolver, key key[T, K], previousKeys []any) T {

	if hasCycle(key, previousKeys) {
		givvPanic("Cycle detected resolving key: %+v", key)
	}

	provider := resolver.providers[key]

	if hasCycleDetection[T](provider) {
		return provider.(providerWithCycleDetection[T]).getWithPreviousKeys(append(previousKeys, key))
	}
	
	return provider.(Provider[T]).Get()
}

func hasCycleDetection[T any](provider any) bool {
	_, ok := provider.(providerWithCycleDetection[T])
	return ok
}

func hasCycle(currentKey any, previousKeys []any) bool {
	for _, key := range previousKeys {
		if key == currentKey {
			return true
		}
	}
	return false
}

func Bind[T any, K any](resolver *Resolver, key key[T, K], provider Provider[T]) {
	resolver.providers[key] = provider
}

func BindToInstance[T any, K any](resolver *Resolver, key key[T, K], value T) {

	if isNil(value) {
		givvPanic("Cannot use BindToInstance() to bind to a nil value. Use BindToNil() instead.")
	}

	Bind(resolver, key, newInstanceProvider(value))
}

func BindToNil[T any, K any](resolver *Resolver, key key[T, K]) {

	typeOf := reflectType[T]()

	if !isNilable(typeOf) {
		givvPanic("BindToNil received type that cannot be nil: %+v", typeOf)
	}

	Bind(resolver, key, newNilProvider[T]())
}

func BindInstanceType[T any](resolver *Resolver, instance T) {
	BindToInstance(resolver, TypeKey[T](), instance)
}

func BindProviderType[T any](resolver *Resolver, provider Provider[T]) {
	Bind(resolver, TypeKey[T](), provider)
}

func BindToFunction[T any, K any](resolver *Resolver, key key[T, K], function func() T) {
	Bind(resolver, key, NewFunctionProvider(function))
}

func BindToFunction1Arg[T any, K any, A any](resolver *Resolver, key key[T, K], function func(A) T, arg ArgSpec[A]) {
	Bind(resolver, key, NewFunction1ArgProvider(resolver, function, arg))
}

func BindToFunction2Args[T any, K any, A1 any, A2 any](
	resolver *Resolver, 
	key key[T, K], 
	function func(A1, A2) T,
	arg1 ArgSpec[A1],
	arg2 ArgSpec[A2],
	) {

	Bind(resolver, key, NewFunction2ArgsProvider(resolver, function, arg1, arg2))
}

func BindToFunction4Args[T any, K any, A1 any, A2 any, A3 any, A4 any](
	resolver *Resolver, 
	key key[T, K], 
	function func(A1, A2, A3, A4) T,
	arg1 ArgSpec[A1],
	arg2 ArgSpec[A2],
	arg3 ArgSpec[A3],
	arg4 ArgSpec[A4],
	) {

	Bind(resolver, key, NewFunction4ArgsProvider(resolver, function, arg1, arg2, arg3, arg4))
}

func BindAutomaticProvider[T any, A any](resolver *Resolver) {
	var provider Provider1Arg[T, A]

	provider = AutomaticProvider[T, A]{
		resolver: resolver,
	}

	BindInstanceType(resolver, provider)
}

func isNil[T any](value T) bool {
	valueOf := reflect.ValueOf(value)
	typeOf := valueOf.Type()
	
	return isNilable(typeOf) && valueOf.IsNil()
}

func isNilable(reflectType reflect.Type) bool {
	switch reflectType.Kind() {
	case reflect.Chan, 
			 reflect.Func, 
			 reflect.Interface, 
			 reflect.Map, 
			 reflect.Pointer, 
			 reflect.Slice:
		return true
	}

	return false
}

type nilProvider[T any] struct{
}

func newNilProvider[T any]() Provider[T] {
	return &nilProvider[T]{}
}

func (provider *nilProvider[T]) Get() T {
	var t T
	return t
}




