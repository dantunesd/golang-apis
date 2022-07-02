package main

import (
	"api-test/fiber/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(handler.Error)
	app.Use(handler.Logger)

	app.Get("/hello-world/:param", handler.HelloWorld)
	app.Post("/test-all/:p1/:p2", handler.TestAll)

	if err := app.Listen(":3000"); err != nil {
		log.Println(err)
	}
}
