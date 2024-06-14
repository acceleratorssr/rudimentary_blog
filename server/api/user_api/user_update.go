package user_api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"server/global"
	"server/models"
	"server/models/res"
	"server/models/stype"
)

type UserUpdate struct {
	ID         uint             `json:"id" binding:"required" msg:"用户id错误"`
	Permission stype.Permission `json:"permission" binding:"oneof=1 2 3 4" msg:"权限参数错误"`
	Username   string           `json:"username"`
	NickName   string           `json:"nick_name"`
	PhoneNum   string           `json:"phone_num"`
}

// 是否需要唯一，false表示不需要唯一
func isOnly(only []string, n string) bool {
	for _, s := range only {
		if s == n {
			return true
		}
	}
	return false
}

// UserUpdate 是一个API视图，用于处理用户更新的请求
//
// @Summary 管理员可强制更新用户信息
// @Description 用户更新视图，需要用户ID，可改权限、用户名、昵称和手机号码。用户名和手机号码需要唯一。
// @Tags 用户
// @Accept json
// @Produce json
// @Param UserUpdate body UserUpdate true "用户ID，可改权限、用户名、昵称和手机号码"
// @Success 200 {string} string "修改成功"
// @Router /api/user_update [put]
func (UserApi) UserUpdate(c *gin.Context) {
	onlyOne := []string{"Username", "PhoneNum"}
	var UUp UserUpdate
	err := c.ShouldBindJSON(&UUp)
	if err != nil {
		global.Log.Warn("UserUpdate -> 参数错误:", err)
		res.FailWithMessage("UserUpdate -> 参数错误", c)
		return
	}

	var user models.UserModels
	err = global.DB.Take(&user, UUp.ID).Error
	if err != nil {
		res.FailWithMessage("UserUpdate -> 用户不存在", c)
		return
	}

	uupType := reflect.TypeOf(UUp)
	uupVal := reflect.ValueOf(UUp)
	if uupType.Kind() == reflect.Ptr {
		uupVal = uupVal.Elem()
		uupType = uupType.Elem()
	}
	for i := 0; i < uupVal.NumField(); i++ {
		f := uupVal.Field(i)
		// 判断字段是否为空
		if !f.IsZero() {
			if isOnly(onlyOne, uupType.Field(i).Name) {
				var u models.UserModels
				global.DB.Where(uupType.Field(i).Name+" = ?", f.Interface()).First(&u)
				if u.ID != user.ID && u.ID != 0 {
					global.Log.Warn("UserUpdate -> 手机号/用户名重复:", uupType.Field(i).Name)
					res.FailWithMessage("UserUpdate -> 手机号/用户名重复", c)
					return
				}
				err = global.DB.Model(&user).Update(uupType.Field(i).Name, f.Interface()).Error
				if err != nil {
					global.Log.Warn("UserUpdate -> 修改失败:", err)
					res.FailWithMessage("UserUpdate -> 修改失败", c)
					return
				}
			} else {
				err = global.DB.Model(&user).Update(uupType.Field(i).Name, f.Interface()).Error
				if err != nil {
					global.Log.Warn("UserUpdate -> 修改失败:", err)
					res.FailWithMessage("UserUpdate -> 修改失败", c)
					return
				}
			}
		}
	}

	//err = global.DB.Model(&user).Update("Permission", UUp.Permission).Error
	//if err != nil {
	//	global.Log.Error("UserUpdate -> 修改权限失败:", err)
	//	res.FailWithMessage("UserUpdate -> 修改权限失败", c)
	//	return
	//}
	res.OKWithAll(UUp, "UserUpdate -> 修改成功:", c)
}
