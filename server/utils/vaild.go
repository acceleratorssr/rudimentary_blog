package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetValidMsg 总体来说，这个函数的目的是从验证错误中提取字段的自定义错误消息（如果存在），否则返回原始错误的字符串表示。这在处理验证错误时，特别是在使用结构体标签定义验证规则时，可以提供更有意义的错误消息。
func GetValidMsg(err error, obj any) string {
	// 使用 reflect 包中的 TypeOf 函数获取传入对象的类型信息，并将其存储在 getObj 变量中
	// 获取的是对象 obj 的指针的反射对象
	getObj := reflect.TypeOf(obj)
	var errs validator.ValidationErrors
	//使用 errors 包中的 As 函数判断 err 是否是 validator.ValidationErrors 类型的错误。如果是，将其转换为 errs 变量
	if errors.As(err, &errs) {
		for _, e := range errs {
			//使用 getObj 中的字段名（从验证错误中获取）尝试获取对象的字段信息。如果字段存在，将其存储在变量 f 中
			// 注意Elem()需要搭配obj此时为指针类型
			// 所以对应的调用这个函数传入的obj必须使用&
			if f, exits := getObj.FieldByName(e.Field()); exits {
				//如果字段存在，通过 Tag.Get("msg") 获取字段的自定义错误消息，并将其作为函数的返回值。
				return f.Tag.Get("msg")
			}
		}
	}
	// 如果传入的错误不是 validator.ValidationErrors 类型，或者无法找到相应的字段信息，那么返回原始错误的字符串表示
	return err.Error()
}
