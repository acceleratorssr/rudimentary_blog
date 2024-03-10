package controller_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/res"
)

type getUsernameRequest struct {
	Username string `form:"username"`
}

func (ControllerApi) GetUsername(c *gin.Context) {
	var data getUsernameRequest
	// 检验，应该从数据库根据对应用户取出密钥进行加密比对，先不加密
	key := c.Request.Header.Get("Key")
	if key != "test" {
		fmt.Println("key error")
		return
	}
	err := c.ShouldBindQuery(&data)
	if err != nil {
		global.Log.Error("GetUsername -> 绑定参数失败" + err.Error())
		res.FailWithMessage("绑定参数失败"+err.Error(), c)
		return
	}

	res.FailWithMessage("你的名字为"+data.Username, c)
}
