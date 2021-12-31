package main

import (
	"github.com/sirupsen/logrus"
)

var (
	NodeRegistry   = map[string]*Node{}
	DefaultRuntime = newRuntime()
)

func main() {
	logrus.Info("Welcome to the Computable PaaS Runtime")

	InitAPI()
}
