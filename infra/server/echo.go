package server

import (
	"service-back/interface/controller"

	"github.com/labstack/echo/v4"
)

type ControllerFunc func(c controller.Context) error

func c(cf ControllerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return cf(c.(controller.Context))
	}
}

type Echo struct {
	*echo.Echo
}

func NewEcho() *Echo {
	return &Echo{Echo: echo.New()}
}

func (e *Echo) EchoGroup(prefix string, m ...echo.MiddlewareFunc) *EchoGroup {
	return &EchoGroup{Group: e.Echo.Group(prefix, m...)}
}

type EchoGroup struct {
	*echo.Group
}

func (eg *EchoGroup) EchoGroup(prefix string, m ...echo.MiddlewareFunc) *EchoGroup {
	return &EchoGroup{Group: eg.Group.Group(prefix, m...)}
}

func (eg *EchoGroup) GET(path string, cf ControllerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return eg.Group.GET(path, c(cf), m...)
}

func (eg *EchoGroup) POST(path string, cf ControllerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return eg.Group.POST(path, c(cf), m...)
}
