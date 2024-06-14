package user_api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"server/global"
	"server/models"
	"server/models/res"
	"server/pkg/utils/jwts"
)

type UserSelf struct {
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	PhoneNum string `json:"phone_num"`
}

// UserSelfUpdate 是一个API视图，用于处理用户自我更新的请求
//
// @Summary 用户自我更新
// @Description 用户自我更新视图，可以更改昵称、头像和手机号码。
// @Tags 用户
// @Accept json
// @Produce json
// @Param UserSelf body UserSelf true "可改昵称，头像，手机号"
// @Success 200 {string} string "修改成功"
// @Router /api/user_self_update [put]
func (UserApi) UserSelfUpdate(c *gin.Context) {
	var user models.UserModels
	var USUp UserSelf

	err := c.ShouldBindJSON(&USUp)
	if err != nil {
		global.Log.Warn("UserSelfUpdate -> 参数错误:", err)
		res.FailWithMessage("UserSelfUpdate -> 参数错误", c)
		return
	}

	_permission, _ := c.Get("parseToken")

	// 注意_permission的类型是 *jwts.Permission
	permission := _permission.(*jwts.CustomClaims)

	// 获取当前登录的用户信息
	err = global.DB.First(&user, "id = ?", permission.UserID).Error
	if err != nil {
		global.Log.Warn("UserSelfUpdate -> 未找到对应用户:", err)
		res.FailWithMessage("UserSelfUpdate -> 未找到对应用户，请检查用户id", c)
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
				global.Log.Warn("UserUpdate -> 修改失败:", err)
				res.FailWithMessage("UserUpdate -> 修改失败", c)
				return
			}
		}
	}
	res.OKWithMessage("修改成功", c)
}
