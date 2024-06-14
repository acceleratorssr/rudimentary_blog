package user_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/pkg/utils/jwts"
	"server/pkg/utils/pwd"
)

// 新旧密码也可以前端校验完，只回传新密码回来
type password struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

// UserChangePassword 是一个API视图，用于处理用户更改密码的请求
//
// @Summary 用户更改密码
// @Description 用户更改密码视图，旧密码验证成功后改为新密码
// @Tags 用户
// @Accept json
// @Produce json
// @Param data body password true "新旧密码"
// @Success 200 {string} string "修改成功"
// @Router /api/user_change_password [put]
func (UserApi) UserChangePassword(c *gin.Context) {
	var user models.UserModels
	var mm password

	err := c.ShouldBindJSON(&mm)
	if err != nil {
		global.Log.Warn("UserChangePassword -> 参数错误:", err)
		res.FailWithMessage("UserChangePassword -> 参数错误", c)
		return
	}

	_permission, _ := c.Get("parseToken")

	// 注意_permission的类型是 *jwts.Permission
	permission := _permission.(*jwts.CustomClaims)

	// 获取当前登录的用户信息
	err = global.DB.First(&user, "id = ?", permission.UserID).Error
	if err != nil {
		global.Log.Warn("UserChangePassword -> 未找到对应用户:", err)
		res.FailWithMessage("UserChangePassword -> 未找到对应用户，请检查用户id", c)
		return
	}

	if pwd.CheckPasswords(user.Password, mm.OldPassword) {
		hashPassword := pwd.HashAndSalt(mm.NewPassword)
		err = global.DB.Model(&user).Update("password", hashPassword).Error
		if err != nil {
			global.Log.Warn("UserChangePassword -> 修改失败:", err)
			res.FailWithMessage("UserChangePassword -> 修改失败", c)
			return
		}
	} else {
		global.Log.Warn("UserSelfUpdate -> 密码错误")
		res.OKWithMessage("密码错误", c)
		return
	}
	res.OKWithMessage("修改成功", c)
}
