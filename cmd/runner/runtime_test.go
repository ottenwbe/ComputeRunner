package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Runtime", func() {
	Context("when program starts", func() {
		It("the default Infrastructure runtime should exist", func() {
			Expect(InfrastructureRuntime).To(Not(BeNil()))
		})
		It("the default Code runtime should exist", func() {
			Expect(CodeRuntime).To(Not(BeNil()))
		})
	})
	Context("when adding a (infrastructure) node", func() {
		Context("with functioning code", func() {
			var (
				nodeCode = "node('ANode', 'abc = 1 + 1; console.log(abc);');"
			)

			It("it should add the node", func() {
				_, err := InfrastructureRuntime.Run(nodeCode)
				Expect(err).To(BeNil())
				Expect(NodeRegistry).To(HaveKey("ANode"))
			})
		})

		Context("with non functioning code", func() {
			var (
				nodeCode = "nod('ANode', 'abc = a  console.log(abc);');"
			)

			It("it should execute the node", func() {
				_, err := InfrastructureRuntime.Run(nodeCode)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
