package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type LoggerMiddleware struct{}

func Logger(c *fiber.Ctx) error {
	var err error
	if err = c.Next(); err != nil {
		log.Println("something went wrong: ", err)
	}
	return err
}
