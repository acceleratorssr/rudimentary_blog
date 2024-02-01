package message_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/utils/jwts"
)

// MessageRequest 目前仅支持发送
type MessageRequest struct {
	ReceiveUserID uint   `json:"receive_user_id" binding:"required" msg:"请输入私聊对象ID"` // 接收者ID
	Content       string `json:"content" binding:"required" msg:"不能发送空消息"`           // 消息内容
}

// MessageSend 是一个API视图，用于处理发送消息的请求
//
// @Summary 发送消息
// @Description 发送消息视图，需要接收者ID和消息内容。已登录的用户可以选择一个用户（包括自己）发送一条消息。
// @Tags 消息
// @Accept json
// @Produce json
// @Param MessageRequest body MessageRequest true "接收者ID&信息内容"
// @Success 200 {string} string "发送成功"
// @Router /api/message_send [post]
func (MessageApi) MessageSend(c *gin.Context) {
	// 已登录的用户，选择一个用户（可以是自己）发送一条消息
	var MR MessageRequest
	var userModel models.UserModels

	err := c.ShouldBindJSON(&MR)
	if err != nil {
		global.Log.Errorln("MessageSend -> 参数绑定失败", err)
		res.FailWithError(err, MessageRequest{}, c)
		return
	}

	_permission, _ := c.Get("parseToken")

	// 注意_permission的类型是 *jwts.Permission
	permission := _permission.(*jwts.CustomClaims)
	// 用户登录后，被删除，token挂掉，应该不用判断发送方的状态？

	// 查看接收方是否存在
	// 但前端如果是微信通讯录那种发起聊天的方式，应该也不用判断用户存在与否？
	err = global.DB.Take(&userModel, MR.ReceiveUserID).Error
	if err != nil {
		global.Log.Errorln("MessageSend -> 查无此用户", err)
		res.FailWithMessage("参数绑定失败", c)
		return
	}

	err = global.DB.Create(&models.MessageModels{
		SendUserID:    permission.UserID,
		ReceiveUserID: MR.ReceiveUserID,
		Content:       MR.Content,
	}).Error
	if err != nil {
		global.Log.Errorln("MessageSend -> 聊天记录保存失败", err)
		res.FailWithMessage("聊天记录发送失败", c)
		return
	}
	res.OKWithMessage("发送成功", c)
}
