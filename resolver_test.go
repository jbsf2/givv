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
		Bind(resolver, key, instance)
		Expect(Resolve(resolver, key)).To(Equal(instance))
	})

	It("can bind a string to an instance", func() {
		instance := &RandomStruct{}
		key := Key[*RandomStruct]("random string")
		Bind(resolver, key, instance)
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
				BindToFunction1Arg(resolver, TypeKey[string](), functionWithOneInputParameter, ArgKey("hello"))
				Bind(resolver, Key[string]("hello"), "hello")
				Expect(Resolve(resolver, TypeKey[string]())).To(Equal(functionWithOneInputParameter("hello")))
			})
		})

		Context("when the function has multiple input params of different types", func() {
			It("resolves each parameter value by looking for a binding for its type", func() {

				street := "731 Market St."
				sf := city{name: "San Francisco"}
				california := state{name: "California"}
				zip := zipcode{code: "94110"}

				Bind(resolver, Key[string]("street"), street)
				Bind(resolver, TypeKey[state](), california)
				Bind(resolver, TypeKey[zipcode](), zip)

				BindToFunction4Args(
					resolver, 
					TypeKey[address](), 
					newAddress,
					ArgKey("street"),
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

				Bind(resolver, TypeKey[RandomInterface](), interfaceImpl)
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

				Bind(resolver, TypeKey[RandomInterface](), pointerToRandomStruct)
				Expect(Resolve(resolver, TypeKey[RandomInterface]())).To(Equal(pointerToRandomStruct))
			})		
		})	
	})

	Describe("BindTypeToInstance()", func() {
		Context("when the value is nil", func() {
			It("panics", func() {
				var nilPointer *RandomStruct
				Expect(func(){BindTypeToInstance(resolver, nilPointer)}).To(Panic())
			})
		})

		Context("when the value is a struct", func() {
			It("successfully binds", func() {
				randomStruct := RandomStruct{}
				BindTypeToInstance(resolver, randomStruct)
				BindToFunction1Arg(
					resolver,
					Key[RandomStruct]("foo"),
					functionWithStructParameter,
					ArgType[RandomStruct](),
				)
				Expect(Resolve(resolver, Key[RandomStruct]("foo"))).To(Equal(functionWithStructParameter(randomStruct)))
			})
		})

		// Context("when the value is a struct pointer", func() {
		// 	It("successfully binds", func() {
		// 		structPointer := &RandomStruct{}
		// 		BindTypeToInstance(resolver, structPointer)
		// 		BindToFunction(resolver, Key[*RandomStruct]("foo"), functionWithStructPointerParameter)
		// 		Expect(Resolve(resolver, Key[*RandomStruct]("foo"))).To(Equal(functionWithStructPointerParameter(structPointer)))
		// 	})
		// })

		// Context("when the value is a channel", func() {
		// 	It("successfully binds", func() {
		// 		channel := make(chan string)
		// 		wrongChannelType := make(chan int)
		// 		BindTypeToInstance(resolver, channel)
		// 		BindTypeToInstance(resolver, wrongChannelType)
		// 		BindToFunction(resolver, Key[chan string]("foo"), functionWithChanParameter)
		// 		Expect(Resolve(resolver, Key[chan string]("foo"))).To(Equal(functionWithChanParameter(channel)))
		// 	})
		// })

		// Context("when the value is a parameterized type", func() {
		// 	It("respects the typing", func() {
		// 		rightMap := 
		// 		wrongChannelType := make(chan int)
		// 		BindTypeToInstance(resolver, channel)
		// 		BindTypeToInstance(resolver, wrongChannelType)
		// 		resolver.BindToFunction("foo", functionWithChanParameter)
		// 		Expect(Resolve(resolver, "foo")).To(Equal(functionWithChanParameter(channel)))
		// 	})
		// })
	})

	// Describe("Binding providers with arguments", func() {
	// 	It("works", func() {
	// 		street := "Guerrero St."
	// 		city := city{name: "San Francisco"}
	// 		california := state{name: "California"}
	// 		zip := zipcode{code: "94110"}

	// 		expectedAddress := newAddress(street, city, california, zip)

	// 		var provider Provider2Args[address, state, zipcode]

	// 		BindProvider2ArgsToFunction4Args(
	// 			resolver, 
	// 			newAddress, 
	// 			city, 
	// 			street,
	// 			DynamicArg[state](),
	// 			DynamicArg[zipcode](),
	// 		)

	// 		provider = Resolve(resolver, TypeKey[Provider2Args[address, state, zipcode]]())

	// 		Expect(provider.Get(california, zip)).To(Equal(expectedAddress))
			
	// 	})
	// })
})
