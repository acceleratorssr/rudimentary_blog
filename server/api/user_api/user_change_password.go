package user_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/utils/jwts"
	"server/utils/pwd"
)

// 新旧密码也可以前端校验完，只回传新密码回来
type password struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

func (UserApi) UserChangePasswordView(c *gin.Context) {
	var user models.UserModels
	var mm password

	err := c.ShouldBindJSON(&mm)
	if err != nil {
		global.Log.Warn("UserChangePasswordView -> 参数错误:", err)
		res.FailWithMessage("UserChangePasswordView -> 参数错误", c)
		return
	}

	_permission, _ := c.Get("parseToken")

	// 注意_permission的类型是 *jwts.Permission
	permission := _permission.(*jwts.CustomClaims)

	// 获取当前登录的用户信息
	err = global.DB.First(&user, "id = ?", permission.UserID).Error
	if err != nil {
		global.Log.Warn("UserChangePasswordView -> 未找到对应用户:", err)
		res.FailWithMessage("UserChangePasswordView -> 未找到对应用户，请检查用户id", c)
		return
	}

	if pwd.CheckPasswords(user.Password, mm.OldPassword) {
		hashPassword := pwd.HashAndSalt(mm.NewPassword)
		err = global.DB.Model(&user).Update("password", hashPassword).Error
		if err != nil {
			global.Log.Warn("UserChangePasswordView -> 修改失败:", err)
			res.FailWithMessage("UserChangePasswordView -> 修改失败", c)
			return
		}
	} else {
		global.Log.Warn("UserSelfUpdateView -> 密码错误")
		res.OKWithMessage("密码错误", c)
		return
	}
	res.OKWithMessage("修改成功", c)
}
