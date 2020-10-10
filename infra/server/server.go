package server

import (
	"service-back/interface/controller"
	"strings"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start() {
	// Echo instance
	e := NewEcho()
	// Middleware
	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: logFormat(),
		}),
		middleware.CORS(),
		middleware.Recover(),
		wrapContext,
	)

	// Controllers
	userController := controller.NewUserController()

	eg := e.EchoGroup("")
	eg.POST("/signup", userController.SignUp)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func logFormat() string {
	format := strings.Replace(middleware.DefaultLoggerConfig.Format, ",", ",\n  ", -1)
	format = strings.Replace(format, "{", "{\n  ", 1)
	format = strings.Replace(format, "}}", "}\n}", 1)
	return format
}
