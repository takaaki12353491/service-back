package controller

import (
	"reflect"
	"service-back/str"
)

func initializeParam(param interface{}) {
	rv := reflect.Indirect(reflect.ValueOf(param))
	for i := 0; i < rv.NumField(); i++ {
		sf := rv.Type().Field(i)
		value := str.ToKebabCase(sf.Name)
		rv.Field(i).SetString(value)
	}
	return
}
