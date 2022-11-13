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

func BindInterface[T any](resolver *Resolver, interfaceType T, value any) {
	if !isInterface(interfaceType) {
		givvPanic("interfaceType: %+v is not an interface type", interfaceType)
	}

	resolver.Bind(reflect.TypeOf(&interfaceType).Elem(), value)
}

func (resolver *Resolver) BindToFunction(key any, function any) {
	resolver.providers[key] = NewFunctionProvider(function, resolver)
}

func isInterface[T any](value T) bool {
	maybeInterfaceType := reflect.TypeOf(&value).Elem()
	return maybeInterfaceType.Kind() == reflect.Interface
}




