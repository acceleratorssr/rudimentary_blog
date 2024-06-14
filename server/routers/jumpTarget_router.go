package routers

import (
	"server/api"
)

func (RG RouterGroup) JumpTargetRouter() {
	jumpTargetApi := api.Groups.JumpTargetApi
	RG.Router.GET("/jumpTarget", jumpTargetApi.JumpTargetList)
	RG.Router.POST("/jumpTarget", jumpTargetApi.JumpTargetCreate)
	RG.Router.PUT("/jumpTarget/:id", jumpTargetApi.JumpTargetUpdate)
	RG.Router.DELETE("/jumpTarget", jumpTargetApi.JumpTargetRemove)
}
