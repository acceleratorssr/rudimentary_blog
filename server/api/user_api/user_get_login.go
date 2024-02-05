package user_api

import (
	"context"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/models/stype"
	"server/utils/jwts"
)

// UserGetLoginView 返回当前已登录用户信息
//
// @Tags 用户
// @Summary  当前用户信息
// @Description 查询当前用户信息
// @Accept  json
// @Router /api/user_get_login [get]
// @Produce json
// @Success 200 {object} res.Response
func (UserApi) UserGetLoginView(c *gin.Context) {
	token := c.Request.Header.Get("token")
	ctx := context.Background()

	if token == "" {
		res.FailWithMessage("请先登录", c)
		return
	}

	// 判断token是否被注销
	keys, _ := global.Redis.Keys(ctx, "token_*").Result()
	for _, key := range keys {
		//fmt.Println("get	", key)
		if key == "token_"+token {
			res.FailWithMessage("token已注销", c)
			return
		}
	}

	parseToken, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMessage("token解析失败", c)
		return
	}

	var userList = models.UserModels{
		Permission: stype.Permission(parseToken.Permissions),
		Token:      token,
	}

	res.OKWithData(userList, c)

	return
}
