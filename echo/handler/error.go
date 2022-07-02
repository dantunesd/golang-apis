package handler

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func Error(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.JSON(500, ErrorResponse{Status: 500, Detail: err.Error()})
		}
		return nil
	}
}
