package givv

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("resolver", func() {
	var resolver *Resolver

	BeforeEach(func() {
		resolver =  NewResolver()
	})

	It("can bind a Type to an instance", func() {
		instance := &RandomStruct{}
		key := TypeKey[*RandomStruct]()
		BindToInstance(resolver, key, instance)
		Expect(Resolve(resolver, key)).To(Equal(instance))
	})

	It("can bind a string to an instance", func() {
		instance := &RandomStruct{}
		key := Key[*RandomStruct]("random string")
		BindToInstance(resolver, key, instance)
		Expect(Resolve(resolver, key)).To(Equal(instance))
	})

	Describe("binding to a function", func() {
		Context("when the function has no input parameters", func() {
			It("invokes the function and returns its return value", func() {
				key := Key[string]("foo")
				BindToFunction(resolver, key, functionWithNoInputParameters)
				Expect(Resolve(resolver, key)).To(Equal(functionWithNoInputParameters()))
			})
		})

		Context("when the function has one input parameter", func() {
			It("resolves the parameter value by looking for a binding for its type", func() {
				BindToFunction1Arg(resolver, TypeKey[string](), functionWithOneInputParameter, ArgKey[string]("hello"))
				BindToInstance(resolver, Key[string]("hello"), "hello")
				Expect(Resolve(resolver, TypeKey[string]())).To(Equal(functionWithOneInputParameter("hello")))
			})
		})

		Context("when the function has multiple input params of different types", func() {
			It("resolves each parameter value by looking for a binding for its type", func() {

				street := "731 Market St."
				sf := city{name: "San Francisco"}
				california := state{name: "California"}
				zip := zipcode{code: "94110"}

				BindToInstance(resolver, Key[string]("street"), street)
				BindToInstance(resolver, TypeKey[state](), california)
				BindToInstance(resolver, TypeKey[zipcode](), zip)

				BindToFunction4Args(
					resolver, 
					TypeKey[address](), 
					newAddress,
					ArgKey[string]("street"),
					ArgValue(sf),
					ArgType[state](),
					ArgType[zipcode](),
				)

				Expect(Resolve(resolver, TypeKey[address]())).To(Equal(newAddress(street, sf, california, zip)))
			})
		})

		Context("when an input parameter has an interface type", func() {
			It("it is able to resolve the interface binding", func() {
				var interfaceImpl RandomInterface
				interfaceImpl = &RandomStruct{}

				BindToInstance(resolver, TypeKey[RandomInterface](), interfaceImpl)
				BindToFunction1Arg(
					resolver, 
					Key[string]("foo"), 
					functionWithInterfaceParameter, 
					ArgType[RandomInterface](),
				)
				Expect(Resolve(resolver, Key[string]("foo"))).To(Equal(functionWithInterfaceParameter(interfaceImpl)))
			})
		})
	})

	Describe("binding to an interface type", func() {
		Context("when the binding value implements the interace", func() {
			It("successfully binds", func() {
		
				var pointerToRandomStruct RandomInterface
				pointerToRandomStruct = &RandomStruct{
					name: "foo",
				}

				BindToInstance(resolver, TypeKey[RandomInterface](), pointerToRandomStruct)
				Expect(Resolve(resolver, TypeKey[RandomInterface]())).To(Equal(pointerToRandomStruct))
			})		
		})	
	})

	Describe("BindTypeToInstance()", func() {
		Context("when the value is nil", func() {
			It("panics", func() {
				var nilPointer *RandomStruct
				Expect(func(){BindInstanceTypeToInstance(resolver, nilPointer)}).To(Panic())
			})
		})

		Context("when the value is a struct", func() {
			It("successfully binds", func() {
				randomStruct := RandomStruct{}
				BindInstanceTypeToInstance(resolver, randomStruct)
				BindToFunction1Arg(
					resolver,
					Key[RandomStruct]("foo"),
					functionWithStructParameter,
					ArgType[RandomStruct](),
				)
				Expect(Resolve(resolver, Key[RandomStruct]("foo"))).To(Equal(functionWithStructParameter(randomStruct)))
			})
		})

		Describe("automatic providers", func() {
			Context("when all parameter types are resolvable", func() {
				It("provides a provider", func() {

					BindToInstance(resolver, Key[string](1), "one")
					BindToInstance(resolver, Key[string](2), "two")

					BindProvider1Arg[string, int](resolver)

					type providerType = Provider1Arg[string, int]
					provider := Resolve(resolver, TypeKey[providerType]())

					Expect(provider.Get(1)).To(Equal("one"))		
					Expect(provider.Get(2)).To(Equal("two"))		
				})
			})
		})
	})
})
