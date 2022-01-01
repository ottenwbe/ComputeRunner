package main

import (
	"github.com/sirupsen/logrus"
)

var (
	NodeRegistry          = map[string]*Node{}
	InfrastructureRuntime = newRuntime("Infrastructure", true)
	CodeRuntime           = newRuntime("Code", false)
)

func main() {
	logrus.Info("Welcome to the Computable PaaS Runtime")

	InitAPI()
}
