package givv

import "reflect"

type key[T any, K any] struct {
	key K
}

func Key[T any, K any](keyValue K) key[T, K] {
	return key[T, K]{
		key: keyValue,
	}
}

func TypeKey[T any]() key[T, reflect.Type] {
	return Key[T](reflectType[T]())
}
