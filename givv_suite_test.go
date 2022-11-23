package givv

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGivv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Givv Suite")
}

type RandomInterface interface {
	DoSomething(string) string
}
type RandomStruct struct {
	name string
}

func (random *RandomStruct) DoSomething(something string) string {
	return "I did " + something
}

type EmptyStruct struct {
}

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

func newAddress(street string, city city, state state, zip zipcode) address {
	return address {
		street: street,
		city: city,
		state: state,
		zip: zip,
	}
}

func functionWithNoInputParameters() string {
	return "hello"
}

func functionWithOneInputParameter(message string) string {
	return message
}

func functionWithThreeInputParameters(message string, city *city, state state) string {
	return message
}

func functionWithInterfaceParameter(randomInterface RandomInterface) string {
	return randomInterface.DoSomething("something")
}

func functionWithStructParameter(randomStruct RandomStruct) RandomStruct {
	return randomStruct
}

func functionWithStructPointerParameter(structPointer *RandomStruct) *RandomStruct {
	return structPointer
}

func functionWithChanParameter(channel chan(string)) chan(string) {
	return channel
}

func functionWithProvider2Arg(street string, city city, provider Provider2Args[address, state, zipcode]) *address {
	return nil
}

type A struct{}
type B struct{}
type C struct{}

func functionAdependsB(b B) A {
	return A{}
}

func functionBdependsC(c C) B {
	return B{}
}

func functionCdependsA(a A) C {
	return C{}
}

