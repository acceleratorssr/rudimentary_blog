package routers

import (
	"server/api"
	"server/pkg/middleware"
)

func (RG RouterGroup) InterfaceRouter() {
	InterfaceApi := api.Groups.InterfaceApi
	RG.Router.GET("/interface_list", middleware.JwtAuthUser(), InterfaceApi.InterfaceList)
	RG.Router.POST("/interface_add", middleware.JwtAuthAdmin(), InterfaceApi.InterfaceAdd)
	RG.Router.POST("/interface_update", middleware.JwtAuthAdmin(), InterfaceApi.InterfaceUpdate)
	RG.Router.POST("/interface_remove", middleware.JwtAuthAdmin(), InterfaceApi.InterfaceRemove)
}
