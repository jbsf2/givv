package givv

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type RandomType struct {}

var _ = Describe("Injector", func() {
	var injector *Injector

	BeforeEach(func() {
		injector =  NewInjector()
	})

	It("can bind an instance to a Type", func() {
		instance := &RandomType{}
		injector.bindInstance(reflect.TypeOf(instance), instance)
		Expect(injector.getInstance(reflect.TypeOf(instance))).To(Equal(instance))
	})

	Describe("middleBits", func() {
		It("returns the 'middle' 32 bits", func() {
		})
  })

	Describe("computeLatency", func() {
		Context("When the ReceptionReport has a LastSenderReport but no Delay", func() {
			It("computes the diff (ms) between 'now' and the LastSenderReport", func() {
			})	
		})
  })
})
