package routers

import (
	"server/api"
	"server/pkg/middleware"
)

func (RG RouterGroup) MessageRouter() {
	messageApi := api.Groups.MessageApi
	RG.Router.POST("/message_send", middleware.JwtAuthUser(), messageApi.MessageSend)
}
