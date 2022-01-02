package main

import (
	"ComputeRunner/pkg/node"
	"errors"
	"github.com/robertkrimen/otto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CodeRequest struct {
	Code string `json:"code"`
}

type Response struct {
	Accepted bool   `json:"accepted"`
	Result   string `json:"result"`
	Error    string `json:"error"`
	Version  string `json:"version"`
	Runtime  string `json:"application"`
}

func NewResponse(accepted bool, result string, error string, runtime string) *Response {
	return &Response{Accepted: accepted, Result: result, Error: error, Runtime: runtime, Version: "1.0"}
}

func InitAPI() {
	r := gin.Default()

	r.POST("/code/run", postCodeToRun)

	r.GET("/infrastructure/nodes", getNodes)
	r.POST("/infrastructure", postInfrastructure)
	r.POST("/infrastructure/:node/run", postRunNode)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		logrus.Errorf("Router exited with Error: %v", err)
	}
}

func postRunNode(c *gin.Context) {
	var result otto.Value
	nodeName := c.Param("node")
	if node, ok := node.NodeRegistry[nodeName]; ok {
		node.Run()
		result = node.WaitForResult()
		c.JSON(http.StatusOK,
			NewResponse(true, result.String(), "", "Code"))
	} else {
		c.JSON(http.StatusNotFound, NewResponse(true, "", errors.New("node not found").Error(), "CodeRuntime"))
	}
}

func postInfrastructure(c *gin.Context) {
	var code CodeRequest
	err := c.Bind(&code)
	if err == nil {
		result, compErr := InfrastructureRuntime.Run(code.Code)
		c.JSON(http.StatusOK, NewResponse(true, result.String(), stringifyError(compErr), InfrastructureRuntime.Name))
	} else {
		logrus.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func getNodes(c *gin.Context) {
	c.JSON(http.StatusOK, node.NodeRegistry)
}

func postCodeToRun(c *gin.Context) {
	var code CodeRequest
	err := c.Bind(&code)
	if err == nil {
		result, compErr := CodeRuntime.Run(code.Code)
		c.JSON(http.StatusOK,
			NewResponse(true, result.String(), stringifyError(compErr), "Code"))
	} else {
		logrus.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func stringifyError(err error) string {
	if err != nil {
		return err.Error()
	} else {
		return ""
	}
}
