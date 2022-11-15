package givv

import (
	"reflect"
)

type FunctionProvider struct {
	function any
	typeOf   reflect.Type
	valueOf  reflect.Value
	resolver *Resolver
	argSpecs []ArgSpec
}

func NewFunctionProvider(resolver *Resolver, function any) *FunctionProvider{
	provider := &FunctionProvider{
		function: function,
		typeOf: reflect.TypeOf(function),
		valueOf: reflect.ValueOf(function),
		resolver: resolver,
		argSpecs: []ArgSpec{},
	}

	argCount := argumentCount(function)
	for index := 0; index < argCount; index++ {
		provider.argSpecs = append(provider.argSpecs, UseArgType())
	}

	return provider
}

func NewFunctionProviderWithArgSpecs(resolver *Resolver, function any, argSpecs []ArgSpec) *FunctionProvider{
	if !isFunction(function) {
		givvPanic("function argument %+v is not actually a function", function)
	}

	if argumentCount(function) < len(argSpecs) {
		givvPanic("too many argument specs: %+v for function: %+v", argSpecs, function)
	}

	if argumentCount(function) > len(argSpecs) {
		givvPanic("too few argument specs: %+v for function: %+v", argSpecs, function)
	}

	return &FunctionProvider{
		function: function,
		typeOf: reflect.TypeOf(function),
		valueOf: reflect.ValueOf(function),
		resolver: resolver,
		argSpecs: argSpecs,
	}
}

func (provider *FunctionProvider) Get() any {
	inValues := provider.inValues()
	// fmt.Printf("inValues: %+v\n", inValues)
	// fmt.Printf("valueOf: %+v", provider.valueOf)
	// fmt.Printf("typeOf: %+v", provider.typeOf)
	return provider.valueOf.Call(inValues)[0].Interface()
}

func (provider *FunctionProvider) inValues() []reflect.Value {
	numIn := provider.typeOf.NumIn()
	inValues := []reflect.Value{}

	for i := 0; i < numIn; i++ {
		inType := provider.typeOf.In(i)
		argSpec := provider.argSpecs[i]
		paramValue := argSpec.resolve(provider.resolver, inType)
		inValues = append(inValues, reflect.ValueOf(paramValue))
	}

	return inValues
}

func (provider *FunctionProvider) isEqual(otherProvider *FunctionProvider) bool {
	argsEqual := reflect.DeepEqual(provider.argSpecs, otherProvider.argSpecs)
	functionsEqual := reflect.ValueOf(provider.function).Pointer() == reflect.ValueOf(otherProvider.function).Pointer()

	return argsEqual && functionsEqual 
}

func isFunction(maybeFunction any) bool {
	typeOf := reflect.TypeOf(maybeFunction)
	
	return typeOf.Kind() == reflect.Func
}

func argumentCount(function any) int {
	typeOf := reflect.TypeOf(function)

	return typeOf.NumIn()
}
