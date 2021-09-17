package core

import (
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

func IsURLSearchParams(value interface{}) (Param, bool) {
	param, ok := value.(Param)
	return param, ok
}
