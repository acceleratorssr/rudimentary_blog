package utils

//
//import (
//	"github.com/gin-gonic/gin"
//	"reflect"
//	"server/global"
//	"server/models/res"
//)
//
//type deal interface {
//	UpdateWithValid()
//}
//
//// 是否需要唯一，false表示不需要唯一
//func isOnly(only []string, n string) bool {
//	for _, s := range only {
//		if s == n {
//			return true
//		}
//	}
//	return false
//}
//
//// UpdateWithValid
//// onlyOne 表示是否只允许一个，true表示只允许一个
//// updateList 为参数绑定后，需要更新的结构体
//// model 对应表对象
//func UpdateWithValid[T any, M any](c *gin.Context, updateList T, model M, onlyOne []string) {
//	uupType := reflect.TypeOf(updateList)
//	uupVal := reflect.ValueOf(updateList)
//	if uupType.Kind() == reflect.Ptr {
//		uupVal = uupVal.Elem()
//		uupType = uupType.Elem()
//	}
//	for i := 0; i < uupVal.NumField(); i++ {
//		f := uupVal.Field(i)
//		// 判断字段是否为空
//		if !f.IsZero() {
//			if isOnly(onlyOne, uupType.Field(i).Name) {
//				var u M
//				uType := reflect.TypeOf(u)
//				uVal := reflect.ValueOf(u)
//				global.DB.Where(uupType.Field(i).Name+" = ?", f.Interface()).First(&u)
//				if u.ID != model.ID && u.ID != 0 {
//					global.Log.Warn("UserUpdateView -> 手机号/用户名重复:", uupType.Field(i).Name)
//					res.FailWithMessage("UserUpdateView -> 手机号/用户名重复", c)
//					return
//				}
//				err := global.DB.Model(&model).Update(uupType.Field(i).Name, f.Interface()).Error
//				if err != nil {
//					global.Log.Warn("UserUpdateView -> 修改失败:", err)
//					res.FailWithMessage("UserUpdateView -> 修改失败", c)
//					return
//				}
//			} else {
//				err := global.DB.Model(&model).Update(uupType.Field(i).Name, f.Interface()).Error
//				if err != nil {
//					global.Log.Warn("UserUpdateView -> 修改失败:", err)
//					res.FailWithMessage("UserUpdateView -> 修改失败", c)
//					return
//				}
//			}
//
//		}
//	}
//}
