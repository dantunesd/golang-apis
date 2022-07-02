package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type LoggerMiddleware struct{}

func (e *LoggerMiddleware) Handle(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		log.Println(c.Errors.String())
		return
	}
}
