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
				notAnInterfaceType := reflect.TypeOf(notAnInterface)

				Expect(func(){BindInterface(resolver, notAnInterfaceType, notAnInterface)}).To(Panic())
			})
		})

		Context("When the value does not implement the interface", func() {
			It("panics", func() {
				var interfaceType RandomInterface
				// BindInterface(resolver, interfaceType, EmptyStruct{})
				Expect(func(){BindInterface(resolver, interfaceType, EmptyStruct{})}).To(Panic())
			})
		})

		Context("When the value implements the interface", func() {
			It("successfully binds", func() {
				var interfaceType RandomInterface
				interfaceImpl := &RandomStruct{}
				// BindInterface(resolver, interfaceType, EmptyStruct{})
				BindInterface(resolver, interfaceType, interfaceImpl)
				resolver.BindToFunction("foo", functionWithInterfaceParameter)
				Expect(resolver.Resolve("foo")).To(Equal(functionWithInterfaceParameter(interfaceImpl)))
			})
		})
	})
})
