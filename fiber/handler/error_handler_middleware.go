package handler

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorHandlerMiddleware struct{}

type ErrorResponse struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func (e *ErrorHandlerMiddleware) Handle(c *fiber.Ctx) error {
	if err := c.Next(); err != nil {
		if e, ok := err.(*fiber.Error); ok {
			return e
		}
		return c.Status(400).JSON(ErrorResponse{Status: 400, Detail: err.Error()})
	}
	return nil
}