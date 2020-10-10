package server

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

func (c *Context) CTX() context.Context {
	return c.Context.Request().Context()
}
