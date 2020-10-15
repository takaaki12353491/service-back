package controller

import (
	"reflect"
	"service-back/str"
)

var pn = newParamName()

type paramName struct {
	Name     string
	Email    string
	Password string
	Identity string
}

func newParamName() *paramName {
	pn := &paramName{}
	rv := reflect.Indirect(reflect.ValueOf(pn))
	for i := 0; i < rv.NumField(); i++ {
		sf := rv.Type().Field(i)
		value := str.ToKebabCase(sf.Name)
		rv.Field(i).SetString(value)
	}
	return pn
}
