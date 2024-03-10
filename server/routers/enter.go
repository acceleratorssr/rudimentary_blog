package routers

import (
	"github.com/gin-contrib/cors"
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
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the world of GoFrame!")
	})

	config := cors.DefaultConfig()
	//config.AllowAllOrigins = true	//全允许
	config.AllowOrigins = []string{"http://localhost:8000", "http://127.0.0.1"}                          // 只允许 "http://127.0.0.1" 这个域名的请求
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}                            // 允许的请求方法
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "authorization", "token"} // 允许的请求头
	//config.AllowCredentials = true                                              // 允许接收和发送cookies
	r.Use(cors.New(config))

	// *any是一个通配符，表示可以匹配/swagger/后的任何内容
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouter := r.Group("api")
	// 创建一个自定义的 RouterGroup 结构体实例，该实例包含了一个具有 "/api" 前缀的路由分组
	routerGroupApp := RouterGroup{apiRouter}
	// 在settings_api.go中定义了SettingsApi结构体，这里使用该结构体初始化路由
	// 该结构体中定义了SettingsInfoView方法，该方法对应路由/settings/info
	routerGroupApp.SettingsRouter()
	routerGroupApp.ImagesRouter()
	routerGroupApp.JumpTargetRouter()
	routerGroupApp.MenuRouter()
	routerGroupApp.UserRouter()
	routerGroupApp.MessageRouter()
	routerGroupApp.InterfaceRouter()
	routerGroupApp.ControllerRouter()

	return r
}
