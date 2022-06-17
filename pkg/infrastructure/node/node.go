package node

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/robertkrimen/otto"
	log "github.com/sirupsen/logrus"

	"ComputeRunner/pkg/application"
)

type Type int

const (
	Asynchronous Type = iota
	Synchronous
)

type Description struct {
	Name       string               `json:"name"`
	Code       string               `json:"code"`
	EntryPoint string               `json:"entrypoint"`
	Type       Type                 `json:"type"`
	Runtime    application.Language `json:"runtime"`
}

type BaseNode struct {
	Description
	ID      uuid.UUID
	result  chan otto.Value
	runtime application.AppRuntime
}

type Node interface {
	Run(input string) (otto.Value, error)
	Node() Description
	WaitForStopped() bool
	WaitForResult() otto.Value
	GetNextResult() (otto.Value, error)
}

func NewNodeFromCode(call otto.FunctionCall) (Node, error) {
	var (
		description *Description
		node        Node
		err         error
	)

	description, err = ExtractNodeStrings(call)
	if err != nil {
		log.Info("node extract error")
		return nil, err
	}

	node, err = NewNode(description)
	if err != nil {
		log.Info("new node error")
		return nil, err
	}

	return node, nil
}

func NewNode(description *Description) (Node, error) {
	var (
		node Node
		err  error
	)

	err = Registry.DoesNodeAlreadyExist(description.Name)
	if err != nil {
		return nil, err
	}

	node = newNode(description)
	addNodeToRegistry(node)

	return node, err
}

func newNode(description *Description) Node {
	return &BaseNode{
		Description: *description,
		ID:          uuid.New(),
		result:      make(chan otto.Value),
		runtime:     runtimeByID(description.Name, description.Runtime),
	}
}

func addNodeToRegistry(node Node) {
	Registry[node.Node().Name] = node
}

func runtimeByID(name string, runtime application.Language) application.AppRuntime {
	return application.NewAppRuntime(fmt.Sprintf("%v_rt", name), runtime)
}

func (n *BaseNode) WaitForStopped() bool {
	<-n.result
	return true
}

func (n *BaseNode) WaitForResult() otto.Value {
	return <-n.result
}

func (n *BaseNode) Node() Description {
	return n.Description
}

func (n *BaseNode) Run(input string) (otto.Value, error) {
	switch n.Type {
	case Asynchronous:
		log.Info("Run Async")
		return n.asyncNodeRun()
	case Synchronous:
		log.Info("Run Sync")
		return n.doRun(input)
	default:
		return n.asyncNodeRun()
	}
}

func (n *BaseNode) asyncNodeRun() (otto.Value, error) {
	go func() {
		_ = n.runtime.BeforeStart(n.EntryPoint)
		value, _ := n.doRun("")

		n.result <- value
	}()

	return otto.UndefinedValue(), nil
}

func (n *BaseNode) GetNextResult() (otto.Value, error) {
	select {
	case x, ok := <-n.result:
		if ok {
			logNode(n).Infof("Value %d was read.\n", x)
		} else {
			logNode(n).Info("Channel closed!")
		}
		return x, nil
	default:
		logNode(n).Info("No value ready, moving on.")
	}
	return otto.UndefinedValue(), nil
}

func (n *BaseNode) doRun(input string) (otto.Value, error) {
	err := n.runtime.BeforeStart(n.EntryPoint)
	if err != nil {
		logNode(n).Errorf("Before Start Scripts ran into errors: %v", err)
	}
	value, err := n.runtime.Run(n.Code)
	if err != nil {
		logNode(n).Errorf("Code ran into errors: %v", err)
	}

	return value, err
}

func logNode(n *BaseNode) *log.Entry {
	return log.WithField("node", n.Name).WithField("node-id", n.ID.String())
}
