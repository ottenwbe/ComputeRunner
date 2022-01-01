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
		node = NewNode(fc, CodeRuntime)

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

	Context("when executed", func() {
		var (
			node *Node
		)

		nameV, _ := otto.ToValue("TestName")
		codeV, _ := otto.ToValue("1 + 1")
		fc := otto.FunctionCall{
			ArgumentList: []otto.Value{nameV, codeV},
		}
		node = NewNode(fc, CodeRuntime)

		It("should return a result", func() {
			node.Run()
			Expect(node.WaitForResult().String()).To(Equal("2"))
		})
		It("should stop the calculation", func() {
			node.Run()
			Expect(node.WaitForStopped()).To(BeTrue())
		})
	})

})
