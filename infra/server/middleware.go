package server

import (
	"github.com/labstack/echo/v4"
)

func wrapContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		context := &Context{Context: c}
		return next(context)
	}
}

func login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func logout(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
