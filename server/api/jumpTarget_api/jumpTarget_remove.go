package jumpTarget_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// JumpTargetRemoveView 删除跳转的目标
// @Tags 跳转的目标
// @Summary  删除跳转目标
// @Description 删除跳转的目标
// @Param data body models.RemoveRequest true "需要删除的id_list"
// @Accept  json
// @Router /api/jumpTarget [delete]
// @Produce json
// @Success 200 {object} res.Response
func (JumpTargetApi) JumpTargetRemoveView(c *gin.Context) {
	var RQ models.RemoveRequest
	err := c.ShouldBindJSON(&RQ)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var jumpTargetList []models.JumpTargetModels
	count := global.DB.Find(&jumpTargetList, RQ.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("对应链接不存在", c)
		return
	}
	// 数据库删除
	global.DB.Delete(&jumpTargetList)
	res.OKWithMessage(fmt.Sprintf("成功删除%d份链接", count), c)
}
