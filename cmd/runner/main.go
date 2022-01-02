package main

import (
	"ComputeRunner/pkg/application"
	"ComputeRunner/pkg/infrastructure"
	"github.com/sirupsen/logrus"
)

var (
	InfrastructureRuntime = infrastructure.NewRuntime("Infrastructure")
	CodeRuntime           = application.NewAppRuntime("Code", application.JAVASCRIPT)
)

func main() {
	logrus.Info("Welcome to the Computable PaaS Runtime")

	InitAPI()
}
