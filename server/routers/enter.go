package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"server/global"
)

// RouterGroup 放入公共的路由组
type RouterGroup struct {
	Router *gin.RouterGroup
}

func InitRouters() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	// *any是一个通配符，表示可以匹配/swagger/后的任何内容
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	router := r.Group("api")

	// 创建一个自定义的 RouterGroup 结构体实例，该实例包含了一个具有 "/api" 前缀的路由分组
	routerGroupApp := RouterGroup{router}
	// 在settings_api.go中定义了SettingsApi结构体，这里使用该结构体初始化路由
	// 该结构体中定义了SettingsInfoView方法，该方法对应路由/settings/info
	routerGroupApp.SettingsRouter()
	routerGroupApp.ImagesRouter()
	routerGroupApp.JumpTargetRouter()

	return r
}
