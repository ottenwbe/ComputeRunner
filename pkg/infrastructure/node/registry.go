package node

import "errors"

type RegistryT map[string]Node

// Registry of all nodes
var Registry = RegistryT{}

// DoesNodeAlreadyExist returns an error when a node with the given name already exists
func (receiver RegistryT) DoesNodeAlreadyExist(name string) error {
	var err error
	if _, ok := receiver[name]; ok {
		err = errors.New("node already exists")
	}
	return err
}
