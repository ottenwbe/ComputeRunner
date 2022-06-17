package account_test

import (
	"ComputeRunner/pkg/account"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	Context("Creation", func() {
		n := account.NewAccount("Test")
		It("should create an object", func() {
			Expect(n).ToNot(BeNil())
		})
		It("should create an uuid", func() {
			Expect(n.ID).ToNot(BeNil())
		})
	})
})
