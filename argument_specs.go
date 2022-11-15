package givv

import "reflect"

type ArgSpec interface {
	resolve(resolver *Resolver, argType reflect.Type) any
}

type argValue struct{
	value any
}

func ArgValue(value any) *argValue{
	return &argValue{value: value}
}

func (argValue *argValue) resolve(resolver *Resolver, argType reflect.Type) any {
	return argValue.value
}

type argKey struct {
	key any
}

func ArgKey(key any) *argKey {
	return &argKey{key: key}
}

func (argKey *argKey) resolve(resolver *Resolver, argType reflect.Type) any {
	return resolver.Resolve(argKey.key)
}

type useArgType struct {
}

func UseArgType() *useArgType {
	return &useArgType{}
}

func (useArgType *useArgType) resolve(resolver *Resolver, argType reflect.Type) any {
	return resolver.Resolve(argType)
}
