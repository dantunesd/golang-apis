package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func HelloWorld(c *gin.Context) {
	if c.Param("param") == "caotic" {
		c.Error(errors.New("THIS IS A CAOTIC HELLO WORLD"))
		return
	}

	c.JSON(200, Response{"hello world :)"})
}
