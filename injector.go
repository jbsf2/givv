package givv


type any interface{}

type InstanceProvider[T any] struct {
	instance T
}

func (instanceProvider *InstanceProvider[T]) Get() T {
	return instanceProvider.instance
}

type Injector struct {
	providers map[any]Provider[any]
}

func NewInjector() *Injector {
	return &Injector{
		providers: map[any]Provider[any]{},
	}
}

func (injector *Injector) GetInstance(key any) any {
	return injector.providers[key].Get()
}

func (injector *Injector) bind(key any, value any) {
	injector.providers[key] = &InstanceProvider[any]{instance: value}
}

func (injector *Injector) bindToFunction(key any, function any) {
	injector.providers[key] = NewFunctionProvider(function, injector)
}




