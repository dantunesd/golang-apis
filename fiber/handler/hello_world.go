package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
}

func HelloWorld(c *fiber.Ctx) error {
	if c.Params("param") == "caotic" {
		return errors.New("THIS IS A CAOTIC HELLO WORLD")
	}

	if c.Params("param") == "error" {
		return fiber.NewError(http.StatusBadRequest)
	}

	return c.Status(200).JSON(Response{"hello world :)"})
}
