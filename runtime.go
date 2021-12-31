package main

import (
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"
)

type Runtime struct {
	vm   *otto.Otto
	Name string
}

func newRuntime() *Runtime {
	tmpRuntime := &Runtime{
		vm:   otto.New(),
		Name: "Default",
	}
	err := tmpRuntime.vm.Set("node", node(tmpRuntime))
	if err != nil {
		logrus.Fatalf("Node method could not be added to runtime")
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
		n := newNode(call)

		go runNode(n, runtime)
		untilStopped(n)

		return otto.Value{}
	}
}

func untilStopped(n *Node) bool {
	return <-n.stop
}

func runNode(n *Node, r *Runtime) {
	_, err := r.vm.Run(n.Code)
	if err != nil {
		logrus.Errorf("Code ran into errors while executing %v", err)
	}
	n.stop <- true
}
