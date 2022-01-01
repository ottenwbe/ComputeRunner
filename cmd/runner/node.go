package main

import (
	"github.com/google/uuid"
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"
)

type Node struct {
	Name    string
	Code    string
	ID      uuid.UUID
	result  chan otto.Value
	runtime *Runtime
}

func NewNode(call otto.FunctionCall, r *Runtime) *Node {
	//TODO: what happens if there are no arguments!
	n := &Node{
		Name:    call.Argument(0).String(),
		Code:    call.Argument(1).String(),
		ID:      uuid.New(),
		result:  make(chan otto.Value),
		runtime: r,
	}
	NodeRegistry[n.Name] = n
	return n
}

func (n *Node) WaitForStopped() bool {
	<-n.result
	return true
}

func (n *Node) WaitForResult() otto.Value {
	return <-n.result
}

func (n *Node) Run() {
	go func() {
		value, err := n.runtime.vm.Run(n.Code)
		if err != nil {
			logrus.Errorf("Code ran into errors while executing %v", err)
		}
		n.result <- value
	}()
}
