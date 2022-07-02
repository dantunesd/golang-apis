package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
}

func HelloWorld(c *fiber.Ctx) error {
	if c.Params("param") == "caotic" {
		return errors.New("THIS IS A CAOTIC HELLO WORLD")
	}

	return c.Status(200).JSON(Response{"hello world :)"})
}
