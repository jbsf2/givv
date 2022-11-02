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

func (provider FunctionProvider) get() any {
	inValues := provider.inValues()
	// fmt.Printf("inValue: %+v\n", inValues)
	return provider.valueOf.Call(inValues)[0].Interface()
}

func (provider FunctionProvider) inValues() []reflect.Value {
	numIn := provider.typeOf.NumIn()
	inValues := []reflect.Value{}

	for i := 0; i < numIn; i++ {
		in := provider.typeOf.In(i)
		instance := provider.injector.getInstance(in)
		inValues = append(inValues, reflect.ValueOf(instance))
	}

	return inValues
}
