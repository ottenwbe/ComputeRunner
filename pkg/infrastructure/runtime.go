package infrastructure

import (
	"github.com/google/uuid"
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"

	"ComputeRunner/pkg/application"
	"ComputeRunner/pkg/node"
)

type Runtime struct {
	vm   *otto.Otto
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func NewRuntime(name string) *Runtime {
	tmpRuntime := &Runtime{
		vm:   otto.New(),
		Name: name,
		ID:   uuid.New(),
	}

	err := tmpRuntime.vm.Set("node", newNode(application.NewAppRuntime("Node", application.JAVASCRIPT)))
	if err != nil {
		logrus.Fatalf("Node method could not be added to application")
	}

	return tmpRuntime
}

func (r *Runtime) Run(code string) (otto.Value, error) {
	logrus.Trace("Running the following code: " + code)
	value, err := r.vm.Run(code)
	logrus.Trace("Got result: " + value.String())
	if err != nil {
		logrus.WithField("Runtime", r.Name).Errorf("Code could not be executed by Runtime: %v", err)
	}
	return value, err
}

func newNode(runtime application.AppRuntime) func(call otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		node.NewNode(call, runtime)
		return otto.Value{}
	}
}
