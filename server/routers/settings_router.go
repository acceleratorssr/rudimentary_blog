package routers

import (
	"server/api"
	"server/middleware"
)

// SettingsRouter 拿到ApiGroups，并调用相应方法
func (RG RouterGroup) SettingsRouter() {
	settingsApi := api.Groups.SettingsApi
	RG.Router.GET("/settings/:name", middleware.JwtAuthAdmin(), settingsApi.SettingsInfoView)
	RG.Router.PUT("/settings", middleware.JwtAuthAdmin(), settingsApi.SettingsUpdate)
}
