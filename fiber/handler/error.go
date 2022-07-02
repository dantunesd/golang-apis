package handler

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func Error(c *fiber.Ctx) error {
	if err := c.Next(); err != nil {
		if e, ok := err.(*fiber.Error); ok { // ignoring fiber errors such as Resource Not Found
			return e
		}
		return c.Status(500).JSON(ErrorResponse{Status: 500, Detail: err.Error()})
	}
	return nil
}
