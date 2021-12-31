package main

import (
	"github.com/google/uuid"
	"github.com/robertkrimen/otto"
)

type Node struct {
	Name    string
	Code    string
	ID      uuid.UUID
	channel chan int
	stop    chan bool
}

func newNode(call otto.FunctionCall) *Node {
	//TODO: what happens if there are no arguments!
	n := &Node{
		Name:    call.Argument(0).String(),
		Code:    call.Argument(1).String(),
		ID:      uuid.New(),
		channel: make(chan int),
		stop:    make(chan bool),
	}
	NodeRegistry[n.Name] = n
	return n
}
