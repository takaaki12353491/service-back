package validator

import (
	"regexp"
	"service-back/errs"

	"gopkg.in/go-playground/validator.v9"
)

var (
	v = newValidator()
)

func newValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		return regexp.MustCompile(`^[a-zA-Z\d]{8,16}$`).MatchString(password)
	})
	return v
}

func Validate(i interface{}) error {
	err := v.Struct(i)
	if err != nil {
		return errs.Invalidated.Wrap(err, err.Error())
	}
	return nil
}
