package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	// TODO:删除用户时，注意用户关联的数据也需要删除
	var RQ models.RemoveRequest
	err := c.ShouldBindJSON(&RQ)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var userList []models.UserModels
	count := global.DB.Find(&userList, RQ.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}
	// 数据库删除
	global.DB.Delete(&userList)
	res.OKWithMessage(fmt.Sprintf("成功删除%d名用户", count), c)
}
