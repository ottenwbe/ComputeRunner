package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CodeRequest struct {
	Code string `json:"code"`
}

func InitAPI() {
	r := gin.Default()
	r.POST("/runtime", func(c *gin.Context) {
		var code CodeRequest
		err := c.Bind(&code)
		if err == nil {
			result, err := DefaultRuntime.Run(code.Code)
			c.JSON(200, gin.H{
				"accepted": true,
				"result":   result.String(),
				"error":    stringifyError(err),
			})
		} else {
			logrus.Error(err)
			c.AbortWithError(400, err)
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func stringifyError(err error) interface{} {
	if err != nil {
		return err.Error()
	} else {
		return ""
	}
}
