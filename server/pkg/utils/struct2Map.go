package utils

import (
	"reflect"
)

// Struct2Map 结构体转map
// tag := valType.Field(i).Tag.Get("json") 不处理tag
// fields为 需要转换的结构体的字段名切片，如果为空，则全部转换
func Struct2Map(st any, fields []string) map[any]any {
	m := make(map[any]any)
	need := make(map[string]bool)

	val := reflect.ValueOf(st)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	valType := val.Type()

	// 判断如果val不是结构体或者val是一个空结构体或者val为空，直接返回空map
	if val.Kind() != reflect.Struct || valType.NumField() == 0 || val.IsZero() {
		return m
	}

	for _, v := range fields {
		need[v] = true
	}

	for i := 0; i < valType.NumField(); i++ {
		name := valType.Field(i).Name
		fieldType := valType.Field(i).Type

		if fieldType.Kind() == reflect.Ptr {
			fieldValue := val.Field(i).Elem()
			m[name] = fieldValue.Interface()
		} else if need[name] || fields == nil {
			m[name] = val.Field(i).Interface()
		}
	}

	return m
}
