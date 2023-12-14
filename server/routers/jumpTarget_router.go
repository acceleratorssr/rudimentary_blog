package routers

import (
	"server/api"
)

func (RG RouterGroup) JumpTargetRouter() {
	jumpTargetApi := api.Groups.JumpTargetApi
	RG.Router.GET("/jumpTarget", jumpTargetApi.JumpTargetListView)
	RG.Router.POST("/jumpTarget", jumpTargetApi.JumpTargetCreateView)
	RG.Router.PUT("/jumpTarget/:id", jumpTargetApi.JumpTargetUpdateView)
	RG.Router.DELETE("/jumpTarget", jumpTargetApi.JumpTargetRemoveView)
}
