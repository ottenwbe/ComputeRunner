package account

import (
	"ComputeRunner/pkg/infrastructure"
	"github.com/google/uuid"
)

// Account details
type Account struct {
	Name  string
	ID    string
	Infra *infrastructure.InfraRuntime
}

// NewAccount is created with a given name
func NewAccount(name string) *Account {
	return &Account{
		ID:    uuid.New().String(),
		Name:  name,
		Infra: infrastructure.NewInfraRuntime(name),
	}
}
