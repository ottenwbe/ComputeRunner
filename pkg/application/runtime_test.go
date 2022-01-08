package application

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Runtime", func() {
	Context("for JavaScript", func() {
		var (
			app = NewAppRuntime("TestJSApp", JAVASCRIPT)
		)
		It("when created it should be able to run javascript", func() {
			result, _ := app.Run("2+2")
			Expect(result.String()).To(Equal("4"))
		})
	})
	Context("Undefined", func() {
		var (
			nilApp = NewAppRuntime("Undefined", -1)
		)

		It("when created it should be nil+", func() {
			Expect(nilApp).ToNot(BeNil())
		})
	})
})
