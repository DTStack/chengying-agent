package utils

import (
	"fmt"
	"reflect"
)

func getAsInterfaceValueAndType(data interface{}) (reflect.Type, reflect.Value, bool) {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Struct {
		data = v.Interface()
	} else if t.Kind() == reflect.Ptr {
		data = v.Elem().Interface()
	} else if t.Kind() != reflect.Interface {
		return reflect.TypeOf(nil), reflect.ValueOf(nil), false
	}
	v = reflect.ValueOf(data)
	t = reflect.TypeOf(data)
	return t, v, true
}

func GetTagValues(data interface{}, tag string) []string {
	list := []string{}
	t, _, ok := getAsInterfaceValueAndType(data)
	if !ok {
		return list
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(tag)
		if tag != "" {
			list = append(list, tag)
		}
	}
	return list
}

func MapTagFromStruct(data interface{}, tag string) (map[string]interface{}, error) {
	t, v, ok := getAsInterfaceValueAndType(data)
	if !ok {
		return nil, fmt.Errorf("Not a struct")
	}
	result := map[string]interface{}{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(tag)
		if tag != "" {
			result[tag] = v.Field(i).Interface()
		}
	}
	return result, nil
}
