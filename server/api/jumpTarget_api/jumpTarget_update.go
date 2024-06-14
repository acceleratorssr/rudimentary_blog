package jumpTarget_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/pkg/utils"
)

// JumpTargetUpdate 更新跳转的目标
//
// @Tags 跳转的目标
// @Summary  更新跳转目标
// @Description 更新跳转的目标
// @Param data body JumpTargetRequest true "需要更新的字段"
// @Accept  json
// @Router /api/jumpTarget [put]
// @Produce json
// @Success 200 {object} res.Response
func (JumpTargetApi) JumpTargetUpdate(c *gin.Context) {
	var jTR JumpTargetRequest
	err := c.ShouldBindJSON(&jTR)
	if err != nil {
		res.FailWithError(err, jTR, c)
		return
	}

	id := c.Param("id")
	var jt models.JumpTargetModels
	err = global.DB.Take(&jt, id).Error
	if err != nil {
		res.FailWithMessage("跳转名称不存在", c)
		return
	}

	// 不用传&，第二个参数为需要转换的字段名，为空默认为全都要
	err = global.DB.Model(&jt).Updates(utils.Struct2Map(jTR, nil)).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改跳转链接失败", c)
		return
	}

	res.OKWithMessage("修改跳转链接成功", c)
}
