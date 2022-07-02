package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type HelloWorld struct{}

type Response struct {
	Message string `json:"message"`
}

func (m *HelloWorld) Handle(c *fiber.Ctx) error {
	if c.Params("param") == "caotic" {
		return errors.New("THIS IS A CAOTIC HELLO WORLD")
	}

	return c.Status(200).JSON(Response{"hello world :)"})
}
