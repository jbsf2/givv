package givv

type AutomaticProvider1Arg[T any, A any] struct {
	resolver *Resolver
}

func (provider AutomaticProvider1Arg[T, A]) Get(arg A) T {
	key := Key[T](arg)

	return Resolve(provider.resolver, key)
}