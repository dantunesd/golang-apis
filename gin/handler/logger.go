package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		log.Println("something went wrong: ", c.Errors.String())
	}
}
