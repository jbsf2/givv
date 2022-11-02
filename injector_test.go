package givv

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type RandomInterface interface {
	DoSomething(string)
}
type RandomStruct struct {
	name string
}

func (random *RandomStruct) DoSomething(str string) {
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
		Expect(injector.getInstance(reflect.TypeOf(instance))).To(Equal(instance))
	})

	It("can bind a string to an instance", func() {
		instance := &RandomStruct{}
		injector.bind("random string", instance)
		Expect(injector.getInstance("random string")).To(Equal(instance))
	})

	Describe("binding to a function", func() {
		Context("when the function has no input parameters", func() {
			It("invokes the function and returns its return value", func() {
				injector.bindToFunction("foo", functionWithNoInputParameters)
				Expect(injector.getInstance("foo")).To(Equal(functionWithNoInputParameters()))
			})
		})

		Context("when the function has one input parameter", func() {
			It("resolves the parameter value by looking for a binding for its type", func() {
				injector.bindToFunction("foo", functionWithOneInputParameter)
				injector.bind(reflect.TypeOf("hello"), "hello")
				Expect(injector.getInstance("foo")).To(Equal(functionWithOneInputParameter("hello")))
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

				Expect(injector.getInstance(addressType)).To(Equal(newAddress(street, city, state, zip)))
			})
		})


	})

	It("can bind an instance to an interface Type ", func() {
		var x RandomInterface

		randomStruct := RandomStruct{
			name: "foo",
		}

		typeOf := reflect.TypeOf(&x).Elem()
		// kind := typeOf.Kind()
		// fmt.Printf("kind: %+v\n", kind)
		// fmt.Printf("typeOf is interface: %v\n", typeOf.Kind() == reflect.Interface)
		// fmt.Printf("typeOf: %+v %T", x, x)

		injector.bind(typeOf, randomStruct)
		Expect(injector.getInstance(typeOf)).To(Equal(randomStruct))
	})

	Describe("computeLatency", func() {
		Context("When the ReceptionReport has a LastSenderReport but no Delay", func() {
			It("computes the diff (ms) between 'now' and the LastSenderReport", func() {
			})	
		})
  })
})
