package interface_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

type removeRequest struct {
	ID uint `json:"id"`
}

// InterfaceRemoveView 是一个API视图，用于处理用户删除的请求
//
// @Summary 接口删除
// @Description 接口删除视图
// @Tags 接口
// @Accept json
// @Produce json
// @Param id_list body object true "接口ID列表"
// @Success 200 {string} string "成功删除%d个接口"
// @Router /api/interface_remove [post]
func (InterfaceApi) InterfaceRemoveView(c *gin.Context) {
	// TODO:删除用户时，注意用户关联的数据也需要删除
	var RQ removeRequest
	err := c.ShouldBindJSON(&RQ)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var userList []models.InterfaceModels
	count := global.DB.Find(&userList, RQ.ID).RowsAffected
	if count == 0 {
		res.FailWithMessage("接口不存在", c)
		return
	}
	// 数据库删除
	global.DB.Delete(&userList)
	res.OKWithMessage(fmt.Sprintf("成功删除%d个接口", count), c)
}
