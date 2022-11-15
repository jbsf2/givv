package givv

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FunctionProvider", func() {
	var resolver *Resolver

	BeforeEach(func() {
		resolver =  NewResolver()
	})

	Describe("NewFunctionProviderWithArgumentKeys", func() {
		Context("when the function isn't a function", func() {
			It("panics", func() {
				notAFunction := city{}
				Expect(func () {NewFunctionProviderWithArgSpecs(resolver, notAFunction, []ArgSpec{})}).To(Panic())
			})
		})

		Context("when there are too many argument keys", func() {
			It("panics", func() {
				function := functionWithOneInputParameter
				argSpecs := []ArgSpec{ArgValue(1), ArgValue(2), ArgValue(3)}
				Expect(func () {NewFunctionProviderWithArgSpecs(resolver, function, argSpecs)}).To(Panic())
			})
		})

		Context("when there are too few argument keys", func() {
			It("panics", func() {
				function := functionWithThreeInputParameters
				Expect(func () {NewFunctionProviderWithArgSpecs(resolver, &function, []ArgSpec{ArgValue(1)})}).To(Panic())
			})
		})

		Context("with valid input params", func() {
			It("returns a FunctionDescriptor", func() {
				function := functionWithThreeInputParameters
				argSpecs := []ArgSpec{ArgValue(1), ArgValue(2), ArgValue(3)}
				expectedValue := &FunctionProvider{
					function: function,
					argSpecs: argSpecs,
				}
				actualValue := NewFunctionProviderWithArgSpecs(resolver, function, argSpecs)
				Expect(actualValue.isEqual(expectedValue)).To(BeTrue())
			})
		})
	})
})
