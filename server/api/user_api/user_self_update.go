package user_api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"server/global"
	"server/models"
	"server/models/res"
	"server/utils/jwts"
)

type UserSelf struct {
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	PhoneNum string `json:"phone_num"`
}

func (UserApi) UserSelfUpdateView(c *gin.Context) {
	var user models.UserModels
	var USUp UserSelf

	err := c.ShouldBindJSON(&USUp)
	if err != nil {
		global.Log.Warn("UserSelfUpdateView -> 参数错误:", err)
		res.FailWithMessage("UserSelfUpdateView -> 参数错误", c)
		return
	}

	_permission, _ := c.Get("parseToken")

	// 注意_permission的类型是 *jwts.Permission
	permission := _permission.(*jwts.CustomClaims)

	// 获取当前登录的用户信息
	err = global.DB.First(&user, "id = ?", permission.UserID).Error
	if err != nil {
		global.Log.Warn("UserSelfUpdateView -> 未找到对应用户:", err)
		res.FailWithMessage("UserSelfUpdateView -> 未找到对应用户，请检查用户id", c)
		return
	}

	// 更新用户信息
	uupType := reflect.TypeOf(USUp)
	uupVal := reflect.ValueOf(USUp)
	if uupType.Kind() == reflect.Ptr {
		uupVal = uupVal.Elem()
		uupType = uupType.Elem()
	}
	for i := 0; i < uupVal.NumField(); i++ {
		f := uupVal.Field(i)
		// 判断字段是否为空
		if !f.IsZero() {
			err = global.DB.Model(&user).Update(uupType.Field(i).Name, f.Interface()).Error
			if err != nil {
				global.Log.Warn("UserUpdateView -> 修改失败:", err)
				res.FailWithMessage("UserUpdateView -> 修改失败", c)
				return
			}
		}
	}
	res.OKWithMessage("修改成功", c)
}
