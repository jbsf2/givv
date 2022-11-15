package givv

import (
	"fmt"
	"reflect"
)

type any interface{}

type InstanceProvider[T any] struct {
	instance T
}

func (instanceProvider *InstanceProvider[T]) Get() T {
	return instanceProvider.instance
}

type Resolver struct {
	providers map[any]Provider[any]
}

func NewResolver() *Resolver {
	return &Resolver{
		providers: map[any]Provider[any]{},
	}
}

func (resolver *Resolver) Resolve(key any) any {
	// Debug("Resolve() key: %+v", key)
	return resolver.providers[key].Get()
}

func (resolver *Resolver) Bind(key any, value any) {
	keyType := reflect.TypeOf(key)
	// Debug("keyType: %+v", keyType)

	var x reflect.Type
	typeInterface := reflect.TypeOf(&x).Elem()

	// Debug("implements: %d", keyType.Implements(typeInterface))
	// Debug("key.Kind(): %+s", keyType.Elem().Kind().String())
	if keyType.Implements(typeInterface) && key.(reflect.Type).Kind() == reflect.Interface {
		// Debug("Yes is interface\n")
		valueType := reflect.TypeOf(value)

		if !valueType.Implements(key.(reflect.Type)) {
			message := fmt.Sprintf("binding key interface %+v not implemented by binding value %+v", key, value)
			panic(message)
		}
	}
 
	// Debug("Bind() key: %+v", key)
	resolver.providers[key] = &InstanceProvider[any]{instance: value}
}

func BindInterface[T any](resolver *Resolver, key T, value T) {
	if !isInterface(key) {
		givvPanic("interfaceType: %+v is not an interface type", key)
	}

	resolver.Bind(reflect.TypeOf(&key).Elem(), value)
}

func BindTypeToInstance[T any](resolver *Resolver, instance T) {
	typeOf := reflect.TypeOf(instance)

	if isNilable(typeOf) && isNil(instance) {
		givvPanic("Cannot use BindTypeToInstance to bind to a nil value")
	}

	resolver.Bind(typeOf, instance)
}

func (resolver *Resolver) BindToFunction(key any, function any) {
	resolver.providers[key] = NewFunctionProvider(resolver, function)
}

func (resolver *Resolver) BindToFunctionWithArgSpecs(key any, function any, argSpecs []ArgSpec) {
	resolver.providers[key] = NewFunctionProviderWithArgSpecs(resolver, function, argSpecs)
}

func isInterface[T any](value T) bool {
	maybeInterfaceType := reflect.TypeOf(&value).Elem()
	return maybeInterfaceType.Kind() == reflect.Interface
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




