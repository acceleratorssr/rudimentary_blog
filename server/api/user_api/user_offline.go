package user_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/res"
	"server/pkg/utils/jwts"
	"time"
)

// Offline 是一个API视图，用于处理用户下线的请求
//
// @Summary 用户下线
// @Description 用户下线视图，直接将当前用户的token注销
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {string} string "success offline"
// @Router /api/user_offline [post]
func (UserApi) Offline(c *gin.Context) {
	ctx := context.Background()
	_permission, _ := c.Get("parseToken")
	permission := _permission.(*jwts.CustomClaims)

	// 从回传的头部中获取token
	token := c.Request.Header.Get("token")

	//time.Duration()
	exp := permission.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)

	// 将token注销掉（值不重要）
	err := global.Redis.Set(ctx, fmt.Sprintf("token_%s", token), "", diff).Err()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("fail offline", c)
		return
	}
	res.OKWithMessage("success offline", c)
}
