package routers

import "server/api"

func (RG RouterGroup) MenuRouter() {
	menuApi := api.Groups.MenuApi
	RG.Router.GET("/menuName", menuApi.MenuNameList)
	RG.Router.GET("/menu", menuApi.MenuList)
	RG.Router.GET("/menuDetail/:id", menuApi.MenuDetail)
	RG.Router.POST("/menu", menuApi.MenuCreate)
	RG.Router.PUT("/menu/:id", menuApi.MenuUpdate)
	RG.Router.DELETE("/menu", menuApi.MenuRemove)
}
