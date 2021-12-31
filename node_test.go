package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/robertkrimen/otto"
)

var _ = Describe("Node", func() {

	Context("when created", func() {

		var (
			node *Node
		)

		nameV, _ := otto.ToValue("TestName")
		codeV, _ := otto.ToValue("abc = 1 + 1; console.log(abc);")
		fc := otto.FunctionCall{
			ArgumentList: []otto.Value{nameV, codeV},
		}
		node = newNode(fc)

		It("should have a given name from the JavaScript code", func() {
			Expect(node.Name).To(Equal("TestName"))
		})
		It("should have a given Code from the JavaScript code", func() {
			Expect(node.Code).To(Equal("abc = 1 + 1; console.log(abc);"))
		})
		It("should be registered in the node NodeRegistry", func() {
			Expect(NodeRegistry).To(HaveKey(node.Name))
		})
	})
})
