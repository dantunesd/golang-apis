package handler

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

func HelloWorld(c echo.Context) error {
	if c.Param("param") == "caotic" {
		return errors.New("THIS IS A CAOTIC HELLO WORLD")
	}

	return c.JSON(200, Response{"hello world :)"})
}
