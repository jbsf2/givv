package givv

type key[T any, K any] struct {
	key K
}

func Key[T any, K any](keyValue K) key[T, K] {
	return key[T, K]{
		key: keyValue,
	}
}

func TypeKey[T any]() key[T, T] {
	var keyValue T

	return Key[T](keyValue)
}