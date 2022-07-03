package main

import (
	"api-test/echo/handler"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(handler.Error)
	app.Use(handler.Logger)

	app.GET("/hello-world/:param", handler.HelloWorld)
	app.POST("/test-all/:p1/:p2/", handler.TestAll)

	if err := app.Start(":3000"); err != nil {
		log.Println(err)
	}
}
