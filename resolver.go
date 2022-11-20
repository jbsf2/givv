package givv

import (
	"reflect"
)

type any interface{}

type key[T any] struct {
	keyValue any
}

func Key[T any](keyValue any) key[T] {
	return key[T]{keyValue: keyValue}
}

func TypeKey[T any]() key[T] {
	var value T

	return key[T]{
		keyValue: reflect.TypeOf(&value).Elem(),
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


func Resolve[T any](resolver *Resolver, key key[T]) T {
	provider := resolver.providers[key]
	if (provider == nil) {
		automaticProvider := automaticProvider(resolver, key)
		if automaticProvider != nil {
			provider = automaticProvider
		}
	}
	return provider.(Provider[T]).Get()
}

func Bind[T any](resolver *Resolver, key key[T], value T) {
	typeOf := reflect.TypeOf(value)

	if isNilable(typeOf) && isNil(value) {
		givvPanic("Cannot use Bind() to bind to a nil value")
	}

	resolver.providers[key] = &InstanceProvider[T]{instance: value}
}

func BindTypeToInstance[T any](resolver *Resolver, instance T) {
	Bind(resolver, TypeKey[T](), instance)
}

func BindTypeToProvider[T any](resolver *Resolver, provider Provider[T]) {
	resolver.providers[TypeKey[T]()] = provider
}

func BindToFunction[T any](resolver *Resolver, key key[T], function func() T) {
	resolver.providers[key] = NewFunctionProvider(function)
}

func BindToFunction1Arg[T any, A any](resolver *Resolver, key key[T], function func(A) T, arg ArgSpec[A]) {
	resolver.providers[key] = NewFunction1ArgProvider(resolver, function, arg)
}

func BindToFunction1ArgErr[T any, A any](resolver *Resolver, key key[T], function func(A) (T, error), arg ArgSpec[A]) {
	resolver.providers[key] = NewFunction1ArgErrProvider(resolver, function, arg)
}

func BindToFunction2Args[T any, A1 any, A2 any](
	resolver *Resolver, 
	key key[T], 
	function func(A1, A2) T,
	arg1 ArgSpec[A1],
	arg2 ArgSpec[A2],
	) {

	resolver.providers[key] = NewFunction2ArgsProvider(resolver, function, arg1, arg2)
}

func BindToFunction4Args[T any, A1 any, A2 any, A3 any, A4 any](
	resolver *Resolver, 
	key key[T], 
	function func(A1, A2, A3, A4) T,
	arg1 ArgSpec[A1],
	arg2 ArgSpec[A2],
	arg3 ArgSpec[A3],
	arg4 ArgSpec[A4],
	) {

	resolver.providers[key] = NewFunction4ArgsProvider(resolver, function, arg1, arg2, arg3, arg4)
}

func automaticProvider[T any](resolver *Resolver, key key[T]) (T, bool) {
	switch key.keyValue {
	// case TypeKey[Provider[T]]():
	// 	return automaticProviderNoArgs[T](resolver), true
	case TypeKey[Provider1Arg[T, A]]():
		return automaticProvider1Arg[T, A](resolver), true
	}

	var x T
	return reflect.TypeOf(&x).Elem().(T), false
}

func automaticProvider1Arg[T any](resolver *Resolver) T {

}


func isNil[T any](value T) bool {
	return reflect.ValueOf(value).IsNil()
}

func isNilable(reflectType reflect.Type) bool {
	switch reflectType.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, 
		reflect.Map, reflect.Pointer, reflect.Slice:
		return true
	}

	return false
}




