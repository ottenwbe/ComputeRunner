package node

import (
	"errors"

	"github.com/google/uuid"
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"

	"ComputeRunner/pkg/application"
)

type Node struct {
	Name    string
	Code    string
	ID      uuid.UUID
	result  chan otto.Value
	runtime application.AppRuntime
}

func NewNode(call otto.FunctionCall, r application.AppRuntime) (*Node, error) {
	var (
		nodeName, nodeCode string
		node               *Node
		err                error
	)

	nodeName, nodeCode, err = extractNodeStrings(call)
	if err != nil {
		return nil, err
	}

	err = doesNodeAlreadyExist(nodeName)
	if err != nil {
		return nil, err
	}

	node = newNode(node, nodeName, nodeCode, r)
	addNodeToRegistry(node)

	return node, err
}

func addNodeToRegistry(node *Node) {
	NodeRegistry[node.Name] = node
}

func newNode(node *Node, nodeName string, nodeCode string, r application.AppRuntime) *Node {
	node = &Node{
		Name:    nodeName,
		Code:    nodeCode,
		ID:      uuid.New(),
		result:  make(chan otto.Value),
		runtime: r,
	}
	return node
}

func extractNodeStrings(call otto.FunctionCall) (nodeName string, nodeCode string, err error) {
	if len(call.ArgumentList) < 2 {
		err = errors.New("not enough arguments in node")
	} else {
		nodeName = call.Argument(0).String()
		nodeCode = call.Argument(1).String()
	}
	return nodeName, nodeCode, err
}

func doesNodeAlreadyExist(name string) error {
	var err error
	if _, ok := NodeRegistry[name]; ok {
		err = errors.New("node already exists")
	}
	return err
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
		value, err := n.runtime.Run(n.Code)
		if err != nil {
			logrus.Errorf("Code ran into errors while executing %v", err)
		}

		n.result <- value
	}()
}
