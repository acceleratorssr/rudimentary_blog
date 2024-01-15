package routers

import (
	"server/api"
	"server/middleware"
)

func (RG RouterGroup) UserRouter() {
	userApi := api.Groups.UserApi
	RG.Router.POST("/user_login", userApi.UsernameLoginView)
	RG.Router.GET("/user_list", middleware.JwtAuthAdmin(), userApi.UserListView)
	RG.Router.PUT("/user_update", middleware.JwtAuthAdmin(), userApi.UserUpdateView)
	RG.Router.PUT("/user_self_update", middleware.JwtAuthUser(), userApi.UserSelfUpdateView)
	RG.Router.PUT("/user_change_password", middleware.JwtAuthUser(), userApi.UserChangePasswordView)
	RG.Router.POST("/user_offline", middleware.JwtAuthUser(), userApi.OfflineView)
	RG.Router.DELETE("/user_remove", middleware.JwtAuthAdmin(), userApi.UserRemoveView)
	RG.Router.PUT("/user_bind_email", middleware.JwtAuthUser(), userApi.UserBindEmailView)
	RG.Router.PUT("/user_register", userApi.UserRegisterView)
}
