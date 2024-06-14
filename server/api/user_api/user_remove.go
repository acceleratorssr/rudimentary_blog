package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// UserRemove 是一个API视图，用于处理用户删除的请求
//
// @Summary 用户删除
// @Description 用户删除视图，需要用户ID列表。删除用户时，注意用户关联的数据也需要删除。
// @Tags 用户
// @Accept json
// @Produce json
// @Param id_list body object true "用户ID列表"
// @Success 200 {string} string "成功删除%d名用户"
// @Router /api/user_remove [delete]
func (UserApi) UserRemove(c *gin.Context) {
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
