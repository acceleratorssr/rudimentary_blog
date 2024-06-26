package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/res"
	"server/models/stype"
	"server/pkg/utils/jwts"
)

func JwtAuth(c *gin.Context) (parseToken *jwts.CustomClaims) {
	token := c.Request.Header.Get("token")
	ctx := context.Background()
	if token == "" {
		global.Log.Error("UserList -> token为空")
		res.FailWithMessage("未登录", c)
		c.Abort()
		return
	}

	parseToken, err := jwts.ParseToken(token)
	if err != nil {
		global.Log.Error("UserList -> token解析失败", err)
		res.FailWithMessage("token解析失败", c)
		c.Abort()
		return
	}

	// 判断token是否被注销
	keys, _ := global.Redis.Keys(ctx, "token_*").Result()
	for _, key := range keys {
		if key == "token_"+token {
			res.FailWithMessage("token已注销", c)
			c.Abort()
			return
		}
	}

	return parseToken
}

func JwtAuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		parseToken := JwtAuth(c)

		if parseToken.Permissions > int(stype.Permission(2)) {
			global.Log.Error("UserList -> 游客权限不足")
			res.FailWithMessage("需要注册后登录进行操作", c)
			c.Abort()
			return
		}
		// 登录的用户
		// Set 是一种将数据存储在当前HTTP请求的上下文中的方法
		// 当前HTTP请求的上下文在请求处理期间将一直存在，并且对于每个请求都是不同的
		// 上下文数据在请求完成后将被删除
		// 在后续处理器函数中，可以使用c.Get函数来获取这个值
		c.Set("parseToken", parseToken)
	}
}

func JwtAuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		parseToken := JwtAuth(c)

		if parseToken.Permissions != int(stype.Permission(1)) {
			global.Log.Error("UserList -> 用户权限不足")
			res.FailWithMessage("用户权限不足", c)
			c.Abort()
			return
		}

		c.Set("parseToken", parseToken)
	}
}
