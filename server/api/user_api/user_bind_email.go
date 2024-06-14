package user_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"server/global"
	"server/models"
	"server/models/res"
	"server/pkg/utils/jwts"
	"server/pkg/utils/pwd"
	"server/pkg/utils/random"
	"time"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱格式出错"`
	Code     *string `json:"code"`
	Password string  `json:"password" binding:"required" msg:"请输入用户密码"`
}

const (
	Code  string = "验证码"
	Note  string = "操作通知"
	Alarm string = "告警通知"
)

type Api struct {
	Subject string
}

// Send 把email_qq配置都注销了，竟然还能正常发送信息
func (a Api) Send(sender, acceptName, body string) error {
	return send(sender, acceptName, a.Subject, body)
}

func NewCode() Api {
	return Api{Subject: Code}
}

func NewNote() Api {
	return Api{Subject: Note}
}

func NewAlarm() Api {
	return Api{Subject: Alarm}
}

func send(sender, name, subject, body string) error {
	e := global.Config.Email163
	return sendMail(
		sender,
		e.UserAddr,
		e.Password,
		e.Host,
		name,
		e.DefaultFromEmail,
		subject,
		body,
		e.Port,
	)
}

func sendMail(sender, userAddr, authCode, host, mailTo, userName, subject, body string, port int) error {
	m := gomail.NewMessage()
	// DefaultFromEmail不能被配置为非邮箱地址
	m.SetHeader("From", m.FormatAddress(userAddr, sender))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	err := d.DialAndSend(m)
	return err
}

// UserBindEmail 是一个处理用户绑定邮箱请求的视图函数;
//
// @Summary 用户绑定邮箱
// @Description 它首先验证用户的存在，然后获取并验证邮箱验证码，如果code为空则代表第一次发送验证码；若验证码验证正确，则它将更新用户的邮箱信息;
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param   code       body    BindEmailRequest     true       "邮箱，用户密码，邮箱验证码"
// @Success 200 {string} string	"返回成功消息"
// @Router /api/user_bind_email [post]
func (UserApi) UserBindEmail(c *gin.Context) {
	var BER BindEmailRequest
	var userModel models.UserModels
	ctx := context.Background()
	_permission, _ := c.Get("parseToken")

	// 注意_permission的类型是 *jwts.Permission
	permission := _permission.(*jwts.CustomClaims)
	err := global.DB.Take(&userModel, "id = ?", permission.UserID).Error
	if err != nil {
		fmt.Println(err)
		res.FailWithMessage("用户不存在", c)
		return
	}

	//1.获取邮箱
	err = c.ShouldBindJSON(&BER)
	if err != nil {
		res.FailWithError(err, BindEmailRequest{}, c)
		return
	}
	if !pwd.CheckPasswords(userModel.Password, BER.Password) {
		global.Log.Warn("绑定邮箱 -> 密码错误")
		res.FailWithMessage("密码错误", c)
		return
	}
	if BER.Code == nil {
		code := random.Code(6)
		err = global.Redis.Set(ctx, userModel.Username, code, time.Minute*10).Err()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("redis出错", c)
			return
		}
		// 只有找到yaml好像就可以了，没找对应qq的都可以
		err = NewCode().Send("acc_1/14", BER.Email, "验证码为 "+code)
		if err != nil {
			fmt.Println(err)
			res.FailWithMessage("邮箱验证码发送失败", c)
			return
		}
		res.OKWithMessage("验证码发送成功，将于10分钟后过期，请及时绑定邮箱", c)
		return
	}
	//2.验证邮箱验证码
	code, e := global.Redis.Get(ctx, userModel.Username).Result()
	if e != nil {
		res.FailWithMessage("验证码已过期", c)
		return
	}
	//3.绑定邮箱
	if *BER.Code == code {
		global.DB.Model(&userModel).Update("email", BER.Email)
	}
	res.OKWithMessage("绑定邮箱成功", c)
}
