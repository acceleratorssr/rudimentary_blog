package stype

import (
	"database/sql/driver"
	"strings"
)

// Array 定义一个自定义类型，类型底层使用[]string。
type Array []string

// Scan 方法的作用是将数据库中的数据（通常是二进制数据）转换为 Go 语言中的类型
// 这里的 Scan 方法接收一个参数 value，其类型是 interface{}。在函数内部
// 首先将 value 转换为字节数组 []byte 类型，并判断其是否为空
// 如果为空，则将 Array 设置为空字符串切片；否则，将字符串按照换行符 \n 分割，然后赋值给 Array
func (a *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) == "" {
		*a = []string{}
		return nil
	}
	*a = strings.Split(string(v), "\n")
	return nil
}

// Value 方法的作用是将 Array 类型转换为数据库可识别的类型，即 driver.Value
// 在这里，它使用 strings.Join 方法将字符串切片中的元素连接起来
// 每个元素之间用换行符 \n 分隔，最终返回一个字符串和一个错误值
func (a Array) Value() (driver.Value, error) {
	return strings.Join(a, "\n"), nil
}
