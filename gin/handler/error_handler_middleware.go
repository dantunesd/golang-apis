package handler

import (
	"github.com/gin-gonic/gin"
)

type ErrorHandlerMiddleware struct{}

type ErrorResponse struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func (e *ErrorHandlerMiddleware) Handle(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(500, ErrorResponse{Status: 500, Detail: c.Errors.String()})
		return
	}
}
