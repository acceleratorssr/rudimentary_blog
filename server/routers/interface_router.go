package routers

import (
	"server/api"
	"server/middleware"
)

func (RG RouterGroup) InterfaceRouter() {
	InterfaceApi := api.Groups.InterfaceApi
	RG.Router.GET("/interface_list", middleware.JwtAuthUser(), InterfaceApi.InterfaceListView)
	RG.Router.POST("/interface_add", middleware.JwtAuthAdmin(), InterfaceApi.InterfaceAddView)
}
