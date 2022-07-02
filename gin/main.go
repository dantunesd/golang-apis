package main

import (
	"api-test/gin/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	errorHandlerMiddleware := handler.ErrorHandlerMiddleware{}
	loggerMiddleware := handler.LoggerMiddleware{}
	helloWorld := handler.HelloWorld{}

	app := gin.Default()

	app.Use(gin.Recovery())
	app.Use(errorHandlerMiddleware.Handle)
	app.Use(loggerMiddleware.Handle)

	app.GET("/hello-world/:param", helloWorld.Handle)

	if err := app.Run(":3000"); err != nil {
		log.Println(err)
	}
}
