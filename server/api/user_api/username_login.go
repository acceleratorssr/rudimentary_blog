package user_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/pkg/utils/jwts"
	"server/pkg/utils/pwd"
)

type UsernameLoginRequest struct {
	Username string `json:"username" binding:"required" msg:"缺少用户名"`
	Password string `json:"password" binding:"required" msg:"缺少密码"`
}

// UsernameLogin 用户登录
// @Summary 用户名登录视图
// @Description 使用用户名和密码进行登录，成功后返回token
// @Tags 用户
// @Accept json
// @Produce json
// @Param ULR body UsernameLoginRequest true "登录请求"
// @Success 200 {object} models.UserModels "成功返回用户信息和token"
// @Router /api/user_login [post]
func (UserApi) UsernameLogin(c *gin.Context) {
	var ULR UsernameLoginRequest
	err := c.ShouldBindJSON(&ULR)
	if err != nil {
		global.Log.Error("UsernameLogin参数验证失败:", err)
		res.FailWithError(err, ULR, c)
		return
	}

	var userModel models.UserModels
	err = global.DB.Take(&userModel, "username = ?", ULR.Username).Error
	if err != nil {
		global.Log.Warn("登录 -> 用户名不存在", err)
		res.FailWithMessage("用户名或密码错误", c)
		return
	}

	if !pwd.CheckPasswords(userModel.Password, ULR.Password) {
		global.Log.Warn("登录 -> 密码错误")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayload{
		Username:    userModel.Username,
		UserID:      userModel.ID,
		Permissions: int(userModel.Permission),
		NickName:    userModel.NickName,
	})
	if err != nil {
		global.Log.Error("token -> 生成失败", err)
		res.FailWithMessage("登录失败", c)
		return
	}
	userModel.Token = token
	res.OKWithData(userModel, c)
}
