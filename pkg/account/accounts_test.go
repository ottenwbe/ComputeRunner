package account_test

import (
	"ComputeRunner/pkg/account"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Accounts", func() {

	Context("add new account", func() {
		err := account.Accounts.Add(account.NewAccount("add1"))
		It("should successfully add it without error", func() {
			Expect(err).To(BeNil())
		})
	})

	Context("add a nil account", func() {
		err := account.Accounts.Add(nil)
		It("should return an error", func() {
			Expect(err).To(Equal(account.NILERROR))
		})
	})

	Context("add a second time an account", func() {
		testName := "a"
		origAccount := account.NewAccount(testName)
		doubleAccount := account.NewAccount(testName)

		err1 := account.Accounts.Add(origAccount)
		It("should be possible", func() {
			Expect(err1).To(BeNil())
		})

		err2 := account.Accounts.Add(doubleAccount)
		It("should not be possible", func() {
			Expect(err2).To(Equal(account.ALREADYEXISTSERROR))
		})
	})

	Context("access by name", func() {

		const testName = "TestGetAndReceive"
		origAccount := account.NewAccount(testName)
		_ = account.Accounts.Add(origAccount)
		resultAccount := account.Accounts.RetrieveByName(testName)

		It("retrieves a valid object", func() {
			Expect(resultAccount).ToNot(BeNil())
		})
		It("returns the correct name", func() {
			Expect(resultAccount.Name).To(Equal(origAccount.Name))
		})
		It("returns the correct infrastructure", func() {
			Expect(resultAccount.Infra).To(Equal(origAccount.Infra))
		})
	})

	Context("access a non existing account by name", func() {

		const testName = "nothere"
		resultAccount := account.Accounts.RetrieveByName(testName)

		It("retrieves a non valid object", func() {
			Expect(resultAccount).To(BeNil())
		})
	})

	Context("access a non existing account by id", func() {

		testName := uuid.New().String()
		resultAccount := account.Accounts.RetrieveByID(testName)

		It("retrieves a non valid object", func() {
			Expect(resultAccount).To(BeNil())
		})
	})

	Context("access by id", func() {

		const testName = "TestGetAndReceiveByID"
		origAccount := account.NewAccount(testName)
		_ = account.Accounts.Add(origAccount)
		resultAccount := account.Accounts.RetrieveByID(origAccount.ID)

		It("retrieves a valid object", func() {
			Expect(resultAccount).ToNot(BeNil())
		})
		It("returns the correct name", func() {
			Expect(resultAccount.Name).To(Equal(origAccount.Name))
		})
	})
})
