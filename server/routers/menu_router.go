package routers

import "server/api"

func (RG RouterGroup) MenuRouter() {
	menuApi := api.Groups.MenuApi
	RG.Router.GET("/menuName", menuApi.MenuNameListView)
	RG.Router.GET("/menu", menuApi.MenuListView)
	RG.Router.POST("/menu", menuApi.MenuCreateView)
}
