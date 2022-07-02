package main

import (
	"api-test/gin/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(gin.Recovery())
	app.Use(handler.Error)
	app.Use(handler.Logger)

	app.GET("/hello-world/:param", handler.HelloWorld)

	if err := app.Run(":3000"); err != nil {
		log.Println(err)
	}
}
