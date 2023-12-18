package routers

import "server/api"

func (RG RouterGroup) MenuRouter() {
	menuApi := api.Groups.MenuApi
	RG.Router.GET("/menuName", menuApi.MenuNameListView)
	RG.Router.GET("/menu", menuApi.MenuListView)
	RG.Router.GET("/menuDetail/:id", menuApi.MenuDetailView)
	RG.Router.POST("/menu", menuApi.MenuCreateView)
	RG.Router.PUT("/menu/:id", menuApi.MenuUpdateView)
	RG.Router.DELETE("/menu", menuApi.MenuRemoveView)
}
