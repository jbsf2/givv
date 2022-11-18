package givv

type ArgSpec[T any] interface {
	resolve(resolver *Resolver) T
	setValue(T)
}

// ----------

type argValue[T any] struct{
	value T
}

func ArgValue[T any](value T) ArgSpec[T]{
	return argValue[T]{value: value}
}

func (argValue argValue[T]) resolve(resolver *Resolver) T {
	return argValue.value
}

func (argValue argValue[T]) setValue(T) {
	// arg value is set at creation time
}

// ----------

type argKey[T any] struct {
	key key[T]
}

func ArgKey[T any](keyValue T) ArgSpec[T] {
	return argKey[T]{key: Key[T](keyValue)}
}

func (argKey argKey[T]) resolve(resolver *Resolver) T {
	return Resolve(resolver, argKey.key)
}

func (argValue argKey[T]) setValue(T) {
}

// ----------

type argType[T any] struct {
	key key[T]
}

func ArgType[T any]() ArgSpec[T] {
	return argType[T]{key: TypeKey[T]()}
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

