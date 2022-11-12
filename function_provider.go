package givv

import (
	"reflect"
)

type FunctionProvider struct {
	function any
	typeOf   reflect.Type
	valueOf  reflect.Value
	resolver *Resolver
}

func NewFunctionProvider(function any, resolver *Resolver) *FunctionProvider{
	return &FunctionProvider{
		function: function,
		typeOf: reflect.TypeOf(function),
		valueOf: reflect.ValueOf(function),
		resolver: resolver,
	}
}

func (provider FunctionProvider) Get() any {
	inValues := provider.inValues()
	// fmt.Printf("inValues: %+v\n", inValues)
	// fmt.Printf("valueOf: %+v", provider.valueOf)
	// fmt.Printf("typeOf: %+v", provider.typeOf)
	return provider.valueOf.Call(inValues)[0].Interface()
}

func (provider FunctionProvider) inValues() []reflect.Value {
	numIn := provider.typeOf.NumIn()
	inValues := []reflect.Value{}

	for i := 0; i < numIn; i++ {
		in := provider.typeOf.In(i)
		// Debug("typeOf in: %+v", in)
		paramValue := provider.resolver.Resolve(in)
		inValues = append(inValues, reflect.ValueOf(paramValue))
	}

	return inValues
}
