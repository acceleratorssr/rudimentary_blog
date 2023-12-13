package routers

import (
	"server/api"
)

// SettingsRouter 拿到ApiGroups，并调用相应方法
func (RG RouterGroup) SettingsRouter() {
	settingsApi := api.Groups.SettingsApi
	RG.Router.GET("/settings/:name", settingsApi.SettingsInfoView)
	RG.Router.PUT("/settings", settingsApi.SettingsUpdate)
}
