package user_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/utils/jwts"
	"server/utils/pwd"
)

type UsernameLoginRequest struct {
	Username string `json:"username" binding:"required" msg:"缺少用户名"`
	Password string `json:"password" binding:"required" msg:"缺少密码"`
}

// UsernameLoginView 用户登录
// @Summary 用户名登录
// @Description 通过用户名和密码进行登录
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param data body UsernameLoginRequest true "用户名及对应密码"
// @Success 200 {string} string	"返回token"
// @Router /api/user_login [post]
func (UserApi) UsernameLoginView(c *gin.Context) {
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
	res.OKWithData(token, c)
}
