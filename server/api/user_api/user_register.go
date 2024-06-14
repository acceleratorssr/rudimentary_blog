package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/models/stype"
	"server/pkg/utils/jwts"
	"server/pkg/utils/pwd"
)

type UserRegisterRequest struct {
	Username string `json:"user_name" msg:"缺少用户名"`
	NickName string `json:"nick_name" msg:"缺少昵称（后续随时可改）"`
	Password string `json:"password"  msg:"缺少密码"`
	Avatar   string `json:"avatar"`
	IP       string `json:"ip"`
	PhoneNum string `json:"phone_num"`
	Email    string `json:"email"`
	// admin:1 user:2 normal:3 banned:4
	//Permission stype.Permission `json:"permission"`
}

// UserRegister 是一个API视图，用于处理用户注册的请求
//
// @Summary 用户注册
// @Description 用户注册视图，需要用户名、昵称和密码。此处前端验证两次输入密码正确后，才会传回信息；会查表以防用户名重复，头像默认，注册成功后自动登录。
// @Tags 用户
// @Accept json
// @Produce json
// @Param UserRegisterRequest body UserRegisterRequest true "用户名，昵称，密码，头像，IP地址，手机号码邮箱"
// @Success 200 {string} string "注册成功"
// @Router /api/user_register [post]
func (UserApi) UserRegister(c *gin.Context) {
	// 注册用户
	var URR UserRegisterRequest
	var userModel models.UserModels
	path := "/uploads/image/1702897857727_zXcxYWSC_2.png"

	err := c.ShouldBindJSON(&URR)
	if err != nil {
		global.Log.Warnln("注册失败 UserRegister -> ", err)
		res.FailWithError(err, UserRegisterRequest{}, c)
		return
	}

	// 注册到mysql，获取id
	err = global.DB.Take(&userModel, "username = ?", URR.Username).Error
	if err == nil {
		global.Log.Warn("注册 -> 用户名已存在", err)
		res.FailWithMessage("用户名已存在", c)
		return
	}

	if URR.Avatar == "" {
		URR.Avatar = path
		fmt.Println("已选择默认头像:", path)
	}

	URR.Password = pwd.HashAndSalt(URR.Password)
	user := models.UserModels{
		Username:       URR.Username,
		NickName:       URR.NickName,
		Password:       URR.Password,
		Avatar:         path,
		Token:          "",
		IP:             "",
		PhoneNum:       URR.PhoneNum,
		SignStatus:     stype.SignNotStatus,
		ArticleModels:  nil,
		CollectsModels: nil,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		res.FailWithMessage("用户名注册失败", c)
		return
	}
	// 返回用户信息
	token, err := jwts.GenToken(jwts.JwtPayload{
		Username:    URR.Username,
		UserID:      user.ID,
		Permissions: int(user.Permission),
		NickName:    URR.NickName,
	})
	if err != nil {
		global.Log.Error("token -> 生成失败", err)
		res.FailWithMessage("登录失败", c)
		return
	}
	res.OKWithData(token, c)
}
