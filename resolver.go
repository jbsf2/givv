package givv

import (
	"reflect"
)

type any interface{}

type key[T any, K any] struct {
	key any
}

func Key[T any, K any](keyValue K) key[T, K] {
	return key[T, K]{
		key: keyValue,
	}
}

func TypeKey[T any]() key[T, T] {
	var keyValue T

	return key[T, T]{
		key: reflect.TypeOf(&keyValue).Elem(),
	}
}

type InstanceProvider[T any] struct {
	instance T
}

func (instanceProvider *InstanceProvider[T]) Get() T {
	return instanceProvider.instance
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
	provider := resolver.providers[key]
	return provider.(Provider[T]).Get()
}

func BindToInstance[T any, K any](resolver *Resolver, key key[T, K], value T) {
	typeOf := reflect.TypeOf(value)

	if isNilable(typeOf) && isNil(value) {
		givvPanic("Cannot use Bind() to bind to a nil value")
	}

	resolver.providers[key] = &InstanceProvider[T]{instance: value}
}

func BindInstanceTypeToInstance[T any](resolver *Resolver, instance T) {
	BindToInstance(resolver, TypeKey[T](), instance)
}

func BindProviderTypeToProvider[T any](resolver *Resolver, provider Provider[T]) {
	resolver.providers[TypeKey[T]()] = provider
}

func BindToFunction[T any, K any](resolver *Resolver, key key[T, K], function func() T) {
	resolver.providers[key] = NewFunctionProvider(function)
}

func BindToFunction1Arg[T any, K any, A any](resolver *Resolver, key key[T, K], function func(A) T, arg ArgSpec[A]) {
	resolver.providers[key] = NewFunction1ArgProvider(resolver, function, arg)
}

func BindToFunction2Args[T any, K any, A1 any, A2 any](
	resolver *Resolver, 
	key key[T, K], 
	function func(A1, A2) T,
	arg1 ArgSpec[A1],
	arg2 ArgSpec[A2],
	) {

	resolver.providers[key] = NewFunction2ArgsProvider(resolver, function, arg1, arg2)
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

	resolver.providers[key] = NewFunction4ArgsProvider(resolver, function, arg1, arg2, arg3, arg4)
}

func BindProvider1Arg[T any, A any](resolver *Resolver) {
	var provider Provider1Arg[T, A]

	provider = AutomaticProvider[T, A]{
		resolver: resolver,
	}

	BindInstanceTypeToInstance(resolver, provider)
}

func isNil[T any](value T) bool {
	return reflect.ValueOf(value).IsNil()
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




