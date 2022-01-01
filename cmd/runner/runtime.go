package main

import (
	"github.com/google/uuid"
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"
)

type Runtime struct {
	vm   *otto.Otto
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func newRuntime(name string, extensions bool) *Runtime {
	tmpRuntime := &Runtime{
		vm:   otto.New(),
		Name: name,
		ID:   uuid.New(),
	}
	if extensions {
		err := tmpRuntime.vm.Set("node", node(newRuntime("Node", false)))
		if err != nil {
			logrus.Fatalf("Node method could not be added to runtime")
		}
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

func node(runtime *Runtime) func(call otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {

		NewNode(call, runtime)

		return otto.Value{}
	}
}
