package givv

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FunctionDescriptor", func() {

	Describe("NewFunctionDescriptor", func() {
		Context("when the function isn't a function", func() {
			It("panics", func() {
				notAFunction := city{}
				Expect(func () {NewFunctionDescriptor(notAFunction, []any{})}).To(Panic())
			})
		})

		Context("when there are too many argument keys", func() {
			It("panics", func() {
				function := functionWithOneInputParameter
				Expect(func () {NewFunctionDescriptor(function, []any{1, 2, 3})}).To(Panic())
			})
		})

		Context("when there are too few argument keys", func() {
			It("panics", func() {
				function := functionWithThreeInputParameters
				Expect(func () {NewFunctionDescriptor(&function, []any{1})}).To(Panic())
			})
		})

		Context("with valid input params", func() {
			It("returns a FunctionDescriptor", func() {
				function := functionWithThreeInputParameters
				argumentKeys := []any{1, 2, 3}
				expectedValue := FunctionDescriptor{
					function: function,
					argumentKeys: argumentKeys,
				}
				actualValue := NewFunctionDescriptor(function, argumentKeys)
				Expect(actualValue.isEqual(&expectedValue)).To(BeTrue())
			})
		})
	})
})
