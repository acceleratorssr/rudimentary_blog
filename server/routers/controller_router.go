package routers

import "server/api"

func (RG RouterGroup) ControllerRouter() {
	ControllerApi := api.Groups.ControllerApi
	RG.Router.GET("/controller", ControllerApi.GetUsername)
}
