package go_axios

import (
	"go-axios/core"
	"reflect"
)

func IsArray(value interface{}) bool {
	typ := reflect.TypeOf(value)
	return typ.Kind() == reflect.Array
}

func IsSlice(value interface{}) bool {
	typ := reflect.TypeOf(value)
	return typ.Kind() == reflect.Slice
}

func IsString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func IsURLSearchParams(value interface{}) (core.Param, bool) {
	param, ok := value.(core.Param)
	return param, ok
}
