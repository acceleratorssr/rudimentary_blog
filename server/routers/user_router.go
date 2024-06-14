package routers

import (
	"server/api"
	"server/pkg/middleware"
)

func (RG RouterGroup) UserRouter() {
	userApi := api.Groups.UserApi
	RG.Router.GET("/user_get_login", userApi.UserGetLogin)
	RG.Router.POST("/user_login", userApi.UsernameLogin)
	RG.Router.GET("/user_list", middleware.JwtAuthAdmin(), userApi.UserList)
	RG.Router.PUT("/user_update", middleware.JwtAuthAdmin(), userApi.UserUpdate)
	RG.Router.PUT("/user_self_update", middleware.JwtAuthUser(), userApi.UserSelfUpdate)
	RG.Router.PUT("/user_change_password", middleware.JwtAuthUser(), userApi.UserChangePassword)
	RG.Router.POST("/user_offline", middleware.JwtAuthUser(), userApi.Offline)
	RG.Router.DELETE("/user_remove", middleware.JwtAuthAdmin(), userApi.UserRemove)
	RG.Router.POST("/user_bind_email", middleware.JwtAuthUser(), userApi.UserBindEmail)
	RG.Router.POST("/user_register", userApi.UserRegister)
}
