package givv

import (
	"reflect"

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
		resolver.Bind(reflect.TypeOf(instance), instance)
		Expect(resolver.Resolve(reflect.TypeOf(instance))).To(Equal(instance))
	})

	It("can bind a string to an instance", func() {
		instance := &RandomStruct{}
		resolver.Bind("random string", instance)
		Expect(resolver.Resolve("random string")).To(Equal(instance))
	})

	Describe("binding to a function", func() {
		Context("when the function has no input parameters", func() {
			It("invokes the function and returns its return value", func() {
				resolver.BindToFunction("foo", functionWithNoInputParameters)
				Expect(resolver.Resolve("foo")).To(Equal(functionWithNoInputParameters()))
			})
		})

		Context("when the function has one input parameter", func() {
			It("resolves the parameter value by looking for a binding for its type", func() {
				resolver.BindToFunction("foo", functionWithOneInputParameter)
				resolver.Bind(reflect.TypeOf("hello"), "hello")
				Expect(resolver.Resolve("foo")).To(Equal(functionWithOneInputParameter("hello")))
			})
		})

		Context("when the function has multiple input params of different types", func() {
			It("resolves each parameter value by looking for a binding for its type", func() {

				street := "731 Market St."
				city := city{name: "San Francisco"}
				state := state{name: "California"}
				zip := zipcode{code: "94110"}

				addressType := reflect.TypeOf(address{})

				resolver.Bind(reflect.TypeOf(street), street)
				resolver.Bind(reflect.TypeOf(city), city)
				resolver.Bind(reflect.TypeOf(state), state)
				resolver.Bind(reflect.TypeOf(zip), zip)

				resolver.BindToFunction(addressType, newAddress)

				Expect(resolver.Resolve(addressType)).To(Equal(newAddress(street, city, state, zip)))
			})
		})

		Context("when an input parameter has an interface type", func() {
			It("it is able to resolve the interface binding", func() {
				var x RandomInterface

				interfaceImpl := RandomStruct{}

				typeOf := reflect.TypeOf(&x).Elem()

				resolver.Bind(typeOf, &interfaceImpl)
				resolver.BindToFunction("foo", functionWithInterfaceParameter)
				Expect(resolver.Resolve("foo")).To(Equal(functionWithInterfaceParameter(&interfaceImpl)))
			})
		})

		Context("with argument specs", func() {
			It("it resolves the arguments using the specs", func() {
				street := "731 Market St."
				city := city{name: "San Francisco"}
				state := state{name: "California"}
				zip := zipcode{code: "94110"}

				addressType := reflect.TypeOf(address{})

				resolver.Bind(reflect.TypeOf(street), street)
				resolver.Bind("city", city)

				argSpecs := []ArgSpec{UseArgType(), ArgKey("city"), ArgValue(state), ArgValue(zip)}

				resolver.BindToFunctionWithArgSpecs(addressType, newAddress, argSpecs)

				Expect(resolver.Resolve(addressType)).To(Equal(newAddress(street, city, state, zip)))
			})
		})
	})

	Describe("binding to an interface type", func() {
		Context("when the binding value implements the interace", func() {
			It("successfully binds", func() {
				var x RandomInterface
				interfaceType := reflect.TypeOf(&x).Elem()
		
				pointerToRandomStruct := &RandomStruct{
					name: "foo",
				}

				structPointerType := reflect.TypeOf(pointerToRandomStruct)

				Expect(structPointerType.Implements(interfaceType)).To(BeTrue())
				
				resolver.Bind(interfaceType, pointerToRandomStruct)
				Expect(resolver.Resolve(interfaceType)).To(Equal(pointerToRandomStruct))
			})		
		})

		Context("when the binding value does not implement the interace", func() {
			It("panics", func() {
				var x RandomInterface
				interfaceType := reflect.TypeOf(&x).Elem()
		
				emptyStruct := EmptyStruct{}
				emptyStructType := reflect.TypeOf(emptyStruct)

				Expect(emptyStructType.Implements(interfaceType)).To(BeFalse())
				
				Expect(func(){resolver.Bind(interfaceType, emptyStruct)}).To(Panic())
			})		
		})		
	})

	Describe("BindInterface()", func() {
		Context("When the key is not a reflect.Type", func() {
			It("panics", func() {
				notAType := "not a type"
				Expect(func(){BindInterface(resolver, notAType, "value")}).To(Panic())
			})
		})

		Context("When the key is not an interface type", func() {
			It("panics", func() {
				notAnInterface := RandomStruct{}
				Expect(func(){BindInterface(resolver, notAnInterface, notAnInterface)}).To(Panic())
			})
		})

		Context("When the value implements the interface", func() {
			It("successfully binds", func() {
				var interfaceImpl RandomInterface
				interfaceImpl = &RandomStruct{}
				// BindInterface(resolver, interfaceType, EmptyStruct{})
				BindInterface(resolver, interfaceImpl, interfaceImpl)
				resolver.BindToFunction("foo", functionWithInterfaceParameter)
				Expect(resolver.Resolve("foo")).To(Equal(functionWithInterfaceParameter(interfaceImpl)))
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
				resolver.BindToFunction("foo", functionWithStructParameter)
				Expect(resolver.Resolve("foo")).To(Equal(functionWithStructParameter(randomStruct)))
			})
		})

		Context("when the value is a struct pointer", func() {
			It("successfully binds", func() {
				structPointer := &RandomStruct{}
				BindTypeToInstance(resolver, structPointer)
				resolver.BindToFunction("foo", functionWithStructPointerParameter)
				Expect(resolver.Resolve("foo")).To(Equal(functionWithStructPointerParameter(structPointer)))
			})
		})

		Context("when the value is a channel", func() {
			It("successfully binds", func() {
				channel := make(chan string)
				wrongChannelType := make(chan int)
				BindTypeToInstance(resolver, channel)
				BindTypeToInstance(resolver, wrongChannelType)
				resolver.BindToFunction("foo", functionWithChanParameter)
				Expect(resolver.Resolve("foo")).To(Equal(functionWithChanParameter(channel)))
			})
		})

		// Context("when the value is a parameterized type", func() {
		// 	It("respects the typing", func() {
		// 		rightMap := 
		// 		wrongChannelType := make(chan int)
		// 		BindTypeToInstance(resolver, channel)
		// 		BindTypeToInstance(resolver, wrongChannelType)
		// 		resolver.BindToFunction("foo", functionWithChanParameter)
		// 		Expect(resolver.Resolve("foo")).To(Equal(functionWithChanParameter(channel)))
		// 	})
		// })
	})
})
