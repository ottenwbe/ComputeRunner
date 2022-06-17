package infrastructure

import (
	"github.com/google/uuid"
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"

	"ComputeRunner/pkg/infrastructure/node"
)

type InfraRuntime struct {
	vm   *otto.Otto
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func NewInfraRuntime(name string) *InfraRuntime {
	tmpRuntime := &InfraRuntime{
		vm:   otto.New(),
		Name: name,
		ID:   uuid.New(),
	}

	err := tmpRuntime.vm.Set("node", newNode())
	if err != nil {
		logrus.Fatalf("Node method could not be added to application")
	}

	return tmpRuntime
}

func (r *InfraRuntime) Run(code string) (otto.Value, error) {
	logrus.Trace("Running Infrastructure Code: " + code)
	value, err := r.vm.Run(code)
	logrus.Trace("Got Infrastructure Code Result: " + value.String())
	if err != nil {
		logrus.WithField("InfraRuntime", r.Name).Errorf("Code could not be executed by InfraRuntime: %v", err)
	}
	return value, err
}

func newNode() func(call otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		_, err := node.NewNodeFromCode(call)
		if err != nil {
			logrus.Error("Error when creating a Node: ", err)
		}
		return otto.Value{}
	}
}
