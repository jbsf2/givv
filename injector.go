package givv

type Provider interface {
	get()interface{}
}

type InstanceProvider struct {
	instance interface{}
}

func (instanceProvider *InstanceProvider) get() interface{} {
	return instanceProvider.instance
}

type Injector struct {
	providers map[interface{}]Provider
}

func NewInjector() *Injector {
	return &Injector{
		providers: map[interface{}]Provider{},
	}
}

func (injector *Injector) getInstance(key interface{}) interface{} {
	return injector.providers[key].get()
}

func (injector *Injector) bindInstance(key interface{}, value interface{}) {
	injector.providers[key] = &InstanceProvider{instance: value}
}




