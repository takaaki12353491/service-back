package server

import (
	"log"
	"service-back/interface/controller"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start() {
	// Echo instance
	e := NewEcho()
	api := e.EchoGroup("")
	// Middleware
	api.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: logFormat(),
		}),
		middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Printf("Request Body: %v\n", string(reqBody))
			log.Printf("Response Body: %v\n", string(resBody))
		}),
		middleware.CORS(),
		middleware.Recover(),
		wrapContext,
	)

	// Controllers
	userController := controller.NewUserController()
	communityController := controller.NewCommunityController()

	// routing
	api.POST("/signup", userController.Signup, logout)
	api.POST("/login", userController.Login, logout)

	communities := api.EchoGroup("/communities")
	communities.GET("", communityController.Index)
	communities.POST("", communityController.Create, login)

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
