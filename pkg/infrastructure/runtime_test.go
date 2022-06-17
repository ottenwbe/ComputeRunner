package infrastructure

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"ComputeRunner/pkg/infrastructure/node"
)

var TestRuntime = NewInfraRuntime("Test")

var _ = Describe("InfraRuntime", func() {
	Context("when program starts", func() {
		It("the default Infrastructure application should exist", func() {
			Expect(TestRuntime).To(Not(BeNil()))
		})
	})
	Context("when adding a (infrastructure) node", func() {
		Context("with functioning code", func() {
			var (
				nodeCode = "node('ANode', 'abc = 1 + 1; console.log(abc);');"
			)

			It("it should add the node", func() {
				_, err := TestRuntime.Run(nodeCode)
				Expect(err).To(BeNil())
				Expect(node.Registry).To(HaveKey("ANode"))
			})
		})

		Context("with non functioning code", func() {
			var (
				nodeCode = "nod('ANode', 'abc = a  console.log(abc);');"
			)

			It("it should execute the node", func() {
				_, err := TestRuntime.Run(nodeCode)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
