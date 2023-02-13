package utility

import (
	"strings"
	"reflect"
)

func GetTypeName(unknown interface{}) string {
	return strings.ToLower(reflect.TypeOf(unknown).String())
}