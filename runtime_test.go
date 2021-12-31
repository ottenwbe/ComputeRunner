package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Runtime", func() {
	Context("when program starts", func() {
		It("the default runtime should exist", func() {
			Expect(DefaultRuntime).To(Not(BeNil()))
		})
	})
	Context("when running a node", func() {
		Context("with functioning code", func() {
			var (
				nodeCode = "node('ANode', 'abc = 1 + 1; console.log(abc);');"
			)

			It("it should execute the node", func() {
				err := DefaultRuntime.Run(nodeCode)
				Expect(err).To(BeNil())
			})
		})

		Context("with non functioning code", func() {
			var (
				nodeCode = "nod('ANode', 'abc = a  console.log(abc);');"
			)

			It("it should execute the node", func() {
				err := DefaultRuntime.Run(nodeCode)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
