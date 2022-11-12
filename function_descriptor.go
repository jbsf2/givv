package givv

import (
	"fmt"
	"reflect"
)

const USE_PARAM_TYPE = "USE_PARAM_TYPE"

type FunctionDescriptor struct {
	function      any
	argumentKeys []any
}

func NewFunctionDescriptor(function any, argumentKeys []any) FunctionDescriptor {

	if !isFunction(function) {
		message := fmt.Sprintf("function argument %+v is not actually a function", function)
		panic(message)
	}

	if argumentCount(function) < len(argumentKeys) {
		message := fmt.Sprintf("too many argument keys: %+v for function: %+v", argumentKeys, function)
		panic(message)
	}

	if argumentCount(function) > len(argumentKeys) {
		message := fmt.Sprintf("too few argument keys: %+v for function: %+v", argumentKeys, function)
		panic(message)
	}

	return FunctionDescriptor{
		function: function,
		argumentKeys: argumentKeys,
	}
}

func (descriptor *FunctionDescriptor) isEqual(otherDescriptor *FunctionDescriptor) bool {
	argsEqual := reflect.DeepEqual(descriptor.argumentKeys, otherDescriptor.argumentKeys)
	functionsEqual := reflect.ValueOf(descriptor.function).Pointer() == reflect.ValueOf(otherDescriptor.function).Pointer()

	return  argsEqual && functionsEqual 
}

func isFunction(maybeFunction any) bool {
	typeOf := reflect.TypeOf(maybeFunction)
	
	return typeOf.Kind() == reflect.Func
}

func argumentCount(function any) int {
	typeOf := reflect.TypeOf(function)

	return typeOf.NumIn()
}