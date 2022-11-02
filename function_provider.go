package givv

import (
	"reflect"
)

type FunctionProvider struct {
	function any
	typeOf   reflect.Type
	valueOf  reflect.Value
	injector *Injector
}

func NewFunctionProvider(function any, injector *Injector) *FunctionProvider{
	return &FunctionProvider{
		function: function,
		typeOf: reflect.TypeOf(function),
		valueOf: reflect.ValueOf(function),
		injector: injector,
	}
}

func (provider FunctionProvider) Get() any {
	inValues := provider.inValues()
	// fmt.Printf("inValue: %+v\n", inValues)
	return provider.valueOf.Call(inValues)[0].Interface()
}

func (provider FunctionProvider) inValues() []reflect.Value {
	numIn := provider.typeOf.NumIn()
	inValues := []reflect.Value{}

	for i := 0; i < numIn; i++ {
		in := provider.typeOf.In(i)
		instance := provider.injector.GetInstance(in)
		inValues = append(inValues, reflect.ValueOf(instance))
	}

	return inValues
}
