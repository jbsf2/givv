package givv

import "reflect"

type ArgSpec[T any] interface {
	resolve(resolver *Resolver) T
	setValue(T)
}

// ----------

type argValue[T any] struct{
	value T
}

func ArgValue[T any](value T) ArgSpec[T]{
	return argValue[T]{
		value: value,
	}
}

func (argValue argValue[T]) resolve(resolver *Resolver) T {
	return argValue.value
}

func (argValue argValue[T]) setValue(T) {
	// arg value is set at creation time
}

// ----------

type argKey[T any, K any] struct {
	key key[T, K]
}

func ArgKey[T any, K any](keyValue K) ArgSpec[T] {
	return argKey[T, K]{
		key: Key[T](keyValue),
	}
}

func (argKey argKey[T, K]) resolve(resolver *Resolver) T {
	return Resolve(resolver, argKey.key)
}

func (argValue argKey[T, K]) setValue(T) {
}

// ----------

type argType[T any] struct {
	key key[T, reflect.Type]
}

func ArgType[T any]() ArgSpec[T] {
	return argType[T]{
		key: TypeKey[T](),
	}
}

func (argType argType[T]) resolve(resolver *Resolver) T {
	return Resolve(resolver, argType.key)
}

func (argType argType[T]) setValue(T) {
}


// -----------

type dynamicArg[T any] struct {
	value T
}

func DynamicArg[T any]() ArgSpec[T] {
	return dynamicArg[T]{}
}

func (arg dynamicArg[T]) resolve(resolver *Resolver) T {
	return arg.value
}

func (arg dynamicArg[T]) setValue(value T) {
	arg.value = value
}

