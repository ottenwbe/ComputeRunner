package application

import (
	"github.com/google/uuid"
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"
)

type Language int

const (
	JAVASCRIPT Language = iota
)

type AppRuntime interface {
	Run(code string) (otto.Value, error)
	BeforeStart(point string) error
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

func (j *JSRuntime) BeforeStart(entryPoint string) error {
	if entryPoint != "" {
		_, err := j.vm.Run(entryPoint)
		return err
	}
	return nil
}

func (j *JSRuntime) Run(code string) (result otto.Value, err error) {
	logrus.Info("Executing: " + code)
	result, err = j.vm.Run(code)
	logrus.Info("res:" + result.String())
	return
}
