package jumpTarget_api

import (
	myutils "github.com/acceleratorssr/My_go_utils"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// JumpTargetUpdateView 更新跳转的目标
// @Tags 跳转的目标
// @Summary  更新跳转目标
// @Description 更新跳转的目标
// @Param data body JumpTargetRequest true "需要更新的字段"
// @Accept  json
// @Router /api/jumpTarget [put]
// @Produce json
// @Success 200 {object} res.Response
func (JumpTargetApi) JumpTargetUpdateView(c *gin.Context) {
	var jTR JumpTargetRequest
	err := c.ShouldBindJSON(&jTR)
	if err != nil {
		res.FailWithError(err, jTR, c)
		return
	}

	id := c.Param("id")
	var jt models.JumpTargetModel
	err = global.DB.Take(&jt, id).Error
	if err != nil {
		res.FailWithMessage("跳转名称不存在", c)
		return
	}

	// 不用传&，第二个参数为不需要转换的字段名，多个字段则直接连着写
	err = global.DB.Model(&jt).Updates(myutils.StructToMap(jTR, "")).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改跳转链接失败", c)
		return
	}

	res.OKWithMessage("修改跳转链接成功", c)
}
