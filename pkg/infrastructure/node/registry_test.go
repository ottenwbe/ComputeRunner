package node

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("The Registry", func() {

	Context("can find duplicates", func() {

		Registry["test"] = newNode(&Description{Name: "test"})

		It("which should result in an error", func() {
			Expect(Registry.DoesNodeAlreadyExist("test")).ToNot(BeNil())
		})
	})

	Context("can identify when a node does not exist", func() {

		It("which should result in no error", func() {
			Expect(Registry.DoesNodeAlreadyExist("testNOTEXISTS")).To(BeNil())
		})
	})

})
