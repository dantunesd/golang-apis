package handler

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		if err = next(c); err != nil {
			log.Println("something went wrong: ", err)
		}
		return err
	}
}
