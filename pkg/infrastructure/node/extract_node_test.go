package node

import (
	"ComputeRunner/pkg/application"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/robertkrimen/otto"
)

var _ = Describe("Extracting Node Strings", func() {

	Context("unsuccessfully with no params", func() {

		var (
			err error
		)

		fc := otto.FunctionCall{
			ArgumentList: []otto.Value{},
		}

		_, err = ExtractNodeStrings(fc)

		It("should result in an error", func() {
			Expect(err).ToNot(BeNil())
		})
	})

	Context("unsuccessfully with wrong format of params", func() {

		var (
			description *Description
		)

		nodeD, _ := otto.ToValue("1+1")
		fc := otto.FunctionCall{
			ArgumentList: []otto.Value{nodeD},
		}

		description, _ = ExtractNodeStrings(fc)

		It("should result in nil description", func() {
			Expect(description).ToNot(BeNil())
		})
	})

	Context("successfully from object", func() {

		var (
			description *Description
			vm          = otto.New()
		)

		vm.Run(`test = {"name":"test", "type":1, code:"2+2" "runtime":0}`)

		nodeD, _ := vm.Get("test") //Object{}({"name": "TestName"})
		fc := otto.FunctionCall{
			ArgumentList: []otto.Value{nodeD},
		}

		description, _ = ExtractNodeStrings(fc)

		It("should return a valid node description", func() {
			Expect(description).ToNot(BeNil())
		})
		It("should have the correct name", func() {
			Expect(description.Name).To(Equal("test"))
		})
		It("should have the correct type", func() {
			Expect(description.Type).To(Equal(Type(1)))
		})
		It("should have the correct runtime", func() {
			Expect(description.Runtime).To(Equal(application.JAVASCRIPT))
		})
		It("should have the correct code", func() {
			Expect(description.Code).To(Equal("2+2"))
		})
	})

	Context("successfully from string", func() {

		var (
			description *Description
		)

		nodeD, _ := otto.ToValue("{\"name\": \"TestName\"}")
		fc := otto.FunctionCall{
			ArgumentList: []otto.Value{nodeD},
		}

		description, _ = ExtractNodeStrings(fc)

		It("should return a valid node description", func() {
			Expect(description).ToNot(BeNil())
		})
		It("should have the correct name", func() {
			Expect(description.Name).To(Equal("TestName"))
		})
	})

})
