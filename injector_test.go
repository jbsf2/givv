package givv

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type RandomInterface interface {
	DoSomething(string) string
}
type RandomStruct struct {
	name string
}

func (random *RandomStruct) DoSomething(something string) string {
	return "I did " + something
}

type OtherRandomStruct struct {}

type state struct {
	name string
	capital city
}

type city struct {
	name string
}

type zipcode struct {
	code string
}

type address struct {
	street string
	city   city
	state  state
	zip    zipcode
}

func functionWithNoInputParameters() string {
	return "hello"
}

func functionWithOneInputParameter(message string) string {
	return message
}

func functionWithInterfaceParameter(randomInterface RandomInterface) string {
	return randomInterface.DoSomething("something")
}

func newAddress(street string, city city, state state, zip zipcode) address {
	return address {
		street: street,
		city: city,
		state: state,
		zip: zip,
	}
}

var _ = Describe("Injector", func() {
	var injector *Injector

	BeforeEach(func() {
		injector =  NewInjector()
	})

	It("can bind a Type to an instance", func() {
		instance := &RandomStruct{}
		injector.bind(reflect.TypeOf(instance), instance)
		Expect(injector.GetInstance(reflect.TypeOf(instance))).To(Equal(instance))
	})

	It("can bind a string to an instance", func() {
		instance := &RandomStruct{}
		injector.bind("random string", instance)
		Expect(injector.GetInstance("random string")).To(Equal(instance))
	})

	Describe("binding to a function", func() {
		Context("when the function has no input parameters", func() {
			It("invokes the function and returns its return value", func() {
				injector.bindToFunction("foo", functionWithNoInputParameters)
				Expect(injector.GetInstance("foo")).To(Equal(functionWithNoInputParameters()))
			})
		})

		Context("when the function has one input parameter", func() {
			It("resolves the parameter value by looking for a binding for its type", func() {
				injector.bindToFunction("foo", functionWithOneInputParameter)
				injector.bind(reflect.TypeOf("hello"), "hello")
				Expect(injector.GetInstance("foo")).To(Equal(functionWithOneInputParameter("hello")))
			})
		})

		Context("when the function has multiple input params of different types", func() {
			It("resolves each parameter value by looking for a binding for its type", func() {

				street := "731 Market St."
				city := city{name: "San Francisco"}
				state := state{name: "California"}
				zip := zipcode{code: "94110"}

				addressType := reflect.TypeOf(address{})

				injector.bind(reflect.TypeOf(street), street)
				injector.bind(reflect.TypeOf(city), city)
				injector.bind(reflect.TypeOf(state), state)
				injector.bind(reflect.TypeOf(zip), zip)

				injector.bindToFunction(addressType, newAddress)

				Expect(injector.GetInstance(addressType)).To(Equal(newAddress(street, city, state, zip)))
			})
		})

		Context("when an input parameter has an interface type", func() {
			It("it is able to resolve the interface binding", func() {
				var x RandomInterface

				interfaceImpl := RandomStruct{}

				typeOf := reflect.TypeOf(&x).Elem()

				injector.bind(typeOf, &interfaceImpl)
				injector.bindToFunction("foo", functionWithInterfaceParameter)
				Expect(injector.GetInstance("foo")).To(Equal(functionWithInterfaceParameter(&interfaceImpl)))
			})
		})
	})

	It("can bind an instance to an interface Type ", func() {
		var x RandomInterface

		randomStruct := RandomStruct{
			name: "foo",
		}

		typeOf := reflect.TypeOf(&x).Elem()

		injector.bind(typeOf, randomStruct)
		Expect(injector.GetInstance(typeOf)).To(Equal(randomStruct))
	})
})
