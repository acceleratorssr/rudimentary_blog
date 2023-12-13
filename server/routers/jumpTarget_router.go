package routers

import (
	"server/api"
)

func (RG RouterGroup) JumpTargetRouter() {
	jumpTargetApi := api.Groups.JumpTargetApi
	RG.Router.POST("/jumpTarget", jumpTargetApi.JumpTargetCreate)
	RG.Router.GET("/jumpTarget", jumpTargetApi.JumpTargetListView)
}
