package main

import (
	"ComputeRunner/pkg/account"
	"ComputeRunner/pkg/infrastructure/node"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertkrimen/otto"
	log "github.com/sirupsen/logrus"
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
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	accPath := r.Group("/accounts")
	accPath.GET("/names/:account", getAccountByName)
	accPath.GET("/:id", getAccountByID)
	accPath.POST("/:account", postAccount)

	r.POST("/application/run", postCodeToRun)

	r.GET("/infrastructure/nodes", getNodes)
	r.POST("/infrastructure", postInfrastructure)
	r.POST("/infrastructure/:node/run", postRunNode)
	r.GET("/infrastructure/:node/result", getNodeExecResult)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Errorf("Router exited with Error: %v", err)
	}
}

func getNodeExecResult(c *gin.Context) {
	nodeName := c.Param("node")
	if node, ok := node.Registry[nodeName]; ok {
		result, err := node.GetNextResult()
		if result.IsUndefined() {
			c.JSON(http.StatusNoContent,
				NewResponse(true, "", stringifyError(err), "Code"))
		} else {
			c.JSON(http.StatusOK,
				NewResponse(true, result.String(), stringifyError(err), "Code"))
		}
	} else {
		c.JSON(http.StatusNotFound, NewResponse(true, "", errors.New("node not found").Error(), "CodeRuntime"))
	}
}

func postRunNode(c *gin.Context) {
	var (
		result otto.Value
		err    error
	)
	nodeName := c.Param("node")
	if node, ok := node.Registry[nodeName]; ok {
		result, err = node.Run("")
		c.JSON(http.StatusOK,
			NewResponse(true, result.String(), stringifyError(err), "Code"))
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
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func getNodes(c *gin.Context) {
	c.JSON(http.StatusOK, node.Registry)
}

func postCodeToRun(c *gin.Context) {
	var code CodeRequest
	err := c.Bind(&code)
	if err == nil {
		result, compErr := CodeRuntime.Run(code.Code)
		c.JSON(http.StatusOK,
			NewResponse(true, result.String(), stringifyError(compErr), "Code"))
	} else {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func postAccount(c *gin.Context) {
	name := c.Param("account")

	acc := account.NewAccount(name)
	err := account.Accounts.Add(acc)

	log.Infof("Account creation for %v", name)

	if err == nil {
		c.JSON(http.StatusCreated, acc)
	} else if err == account.ALREADYEXISTSERROR {
		c.AbortWithError(http.StatusConflict, err)
	} else if err == account.NILERROR {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func getAccountByName(c *gin.Context) {
	name := c.Param("account")

	acc := account.Accounts.RetrieveByName(name)

	if acc != nil {
		c.JSON(http.StatusOK, acc)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}

}

func getAccountByID(c *gin.Context) {
	id := c.Param("id")

	acc := account.Accounts.RetrieveByID(id)

	if acc != nil {
		c.JSON(http.StatusOK, acc)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}

}

func stringifyError(err error) string {
	if err != nil {
		return err.Error()
	} else {
		return ""
	}
}
