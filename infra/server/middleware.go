package server

import (
	"fmt"
	"net/http"
	"os"
	"service-back/consts"
	"service-back/errs"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func wrapContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		context := &Context{Context: c}
		return next(context)
	}
}

func login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(echo.HeaderAuthorization)
		jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv(consts.SIGNINKEY)), nil
		})
		if err != nil {
			log.Error(err)
			c.String(http.StatusUnauthorized, "Login required")
			return err
		}
		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			err = fmt.Errorf("not found claims in %s", tokenString)
			c.String(http.StatusUnauthorized, "Login required")
			return err
		}
		userID, ok := claims["sub"].(string)
		if !ok {
			err = fmt.Errorf("not found %s in %s", "sub", tokenString)
			c.String(http.StatusUnauthorized, "Login required")
			return err
		}
		c.Set(userIDKey, userID)
		return next(c)
	}
}

func logout(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(echo.HeaderAuthorization)
		jwtToken, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if jwtToken != nil {
			errMsg := "Logout required"
			log.Error(errMsg)
			c.String(http.StatusBadRequest, errMsg)
			return errs.Invalidated.New(errMsg)
		}
		return next(c)
	}
}
