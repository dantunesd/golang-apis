package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type HelloWorld struct{}

type Response struct {
	Message string `json:"message"`
}

func (m *HelloWorld) Handle(c *gin.Context) {
	if c.Param("param") == "caotic" {
		c.Error(errors.New("THIS IS A CAOTIC HELLO WORLD"))
		return
	}

	c.JSON(200, Response{"hello world :)"})
}
