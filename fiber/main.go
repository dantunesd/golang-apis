package main

import (
	"api-test/fiber/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	errorHandlerMiddleware := handler.ErrorHandlerMiddleware{}
	loggerMiddleware := handler.LoggerMiddleware{}
	helloWorld := handler.HelloWorld{}

	app := fiber.New()

	app.Use(recover.New())
	app.Use(errorHandlerMiddleware.Handle)
	app.Use(loggerMiddleware.Handle)

	app.Get("/hello-world/:param", helloWorld.Handle)

	if err := app.Listen(":3000"); err != nil {
		log.Println(err)
	}
}
