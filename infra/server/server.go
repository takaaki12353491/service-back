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

	// Middleware
	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: logFormat(),
		}),
		middleware.CORS(),
		middleware.Recover(),
	)

	// Controllers
	userController := controller.NewUserController()
	communityController := controller.NewCommunityController()

	// api
	api := e.EchoGroup("")
	api.Use(
		middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Printf("Request Body: %v\n", string(reqBody))
			log.Printf("Response Body: %v\n", string(resBody))
		}),
		wrapContext,
	)
	// routing
	api.POST("/signup", userController.Signup, logout)
	api.POST("/login", userController.Login, logout)

	communities := api.EchoGroup("/communities")
	communities.GET("", communityController.Index)
	communities.GET("/:id", communityController.Show)
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
