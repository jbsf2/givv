package givv

import (
	"fmt"
	"reflect"
)

type any interface{}

func givvPanic(messageFormat string, args... any) {
	message := fmt.Sprintf(messageFormat, args)
	panic(message)
}

func reflectType[T any]() reflect.Type {
	var value T

	typeOf := reflect.TypeOf(value)
	if typeOf == nil {
		typeOf = reflect.TypeOf(&value).Elem()
	}

	return typeOf
}