package utils

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"reflect"
)

// DeleteValueFromStruct filters struct properties
// Solution taken from Sarath Sadasivan Pillai => https://sarathsp.com/
// Modified for our needs
func DeleteValueFromStruct(key string, object interface{}) error {
	v := reflect.ValueOf(object)
	camelizedKey := strcase.ToCamel(key)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// we only accept structs
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("Only accepts structs; got %T", v)
	}
	v = v.FieldByName(camelizedKey)
	v.Set(reflect.Zero(v.Type()))

	return nil
}
