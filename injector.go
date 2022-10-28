package givv

import (
	"reflect"
)


type any interface{}

type Provider[T any] interface {
	get() T
}

type InstanceProvider[T any] struct {
	instance T
}

func (instanceProvider *InstanceProvider[T]) get() T {
	return instanceProvider.instance
}

type Injector struct {
	providers map[any]Provider[any]
}

func NewInjector() *Injector {
	return &Injector{
		providers: map[any]Provider[any]{},
	}
}

func (injector *Injector) getInstance(key any) any {
	return injector.providers[key].get()
}

func (injector *Injector) bind(key any, value any) {
	injector.providers[key] = &InstanceProvider[any]{instance: value}
}

func (injector *Injector) bindToFunction(key any, function any) {


	// typeOf := reflect.TypeOf(function)

	// in := typeOf.In(0)
	// out := typeOf.Out(0)
	// inType := reflect.TypeOf(in)
	// inKind := in.Kind()
	// inValue := reflect.ValueOf(inType)

	// fmt.Printf("foo: %+v\n", fooType)
	// fmt.Printf("in: %+v\n", in)
	// fmt.Printf("inType: %+v\n", inType)
	// fmt.Printf("inKind: %+v\n", inKind)
	// fmt.Printf("inValue: %+v\n", inValue)
	// fmt.Printf("in nameOff: %+v\n", in.String())
	// fmt.Printf("type nameOff: %+v\n", inType.String())
	// fmt.Printf("value nameOff: %+v\n", inValue.String())

	// fmt.Printf("type: %+v\n", typeOf)
	// fmt.Printf("kind: %+v\n", typeOf.Kind())
	// fmt.Printf("in: %+v\n", in)
	// fmt.Printf("out: %+v\n", out)


	injector.providers[key] = FunctionProvider{
		function: function,
		typeOf: reflect.TypeOf(function),
		valueOf: reflect.ValueOf(function),
		injector: injector,
	}
}




