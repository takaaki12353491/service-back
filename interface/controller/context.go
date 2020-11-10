package controller

import (
	"context"
	"mime/multipart"
)

type Context interface {
	UserID() string
	CTX() context.Context
	String(code int, s string) error
	JSON(code int, i interface{}) error
	Redirect(code int, url string) error
	Bind(i interface{}) error
	FormValue(name string) string
	FormFile(name string) (*multipart.FileHeader, error)
	QueryParam(name string) string
	Param(name string) string
}
