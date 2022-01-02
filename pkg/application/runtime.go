package application

import (
	"github.com/google/uuid"
	"github.com/robertkrimen/otto"
)

type Language int

const (
	JAVASCRIPT Language = iota
)

type AppRuntime interface {
	Run(code string) (otto.Value, error)
}

func NewAppRuntime(name string, language Language) AppRuntime {
	switch language {
	case JAVASCRIPT:
		return &JSRuntime{
			vm:   otto.New(),
			Name: name,
			ID:   uuid.New(),
		}
	default:
		return nil
	}
}

type JSRuntime struct {
	vm   *otto.Otto
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func (j *JSRuntime) Run(code string) (result otto.Value, err error) {
	result, err = j.vm.Run(code)
	return
}
