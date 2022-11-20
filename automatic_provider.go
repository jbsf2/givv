package givv

type AutomaticProvider[T any, A any] struct {
	resolver *Resolver
}

func (provider AutomaticProvider[T, A]) Get(arg A) T {
	key := Key[T](arg)

	return Resolve(provider.resolver, key)
}
