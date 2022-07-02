package handler

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func Error(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(500, ErrorResponse{Status: 500, Detail: c.Errors.String()})
		return
	}
}
