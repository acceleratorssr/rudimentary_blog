package jumpTarget_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

func (JumpTargetApi) JumpTargetUpdateView(c *gin.Context) {
	var JTR JumpTargetRequest
	err := c.ShouldBindJSON(&JTR)
	if err != nil {
		res.FailWithError(err, JTR, c)
		return
	}

	id := c.Param("id")
	var jt models.JumpTargetModel
	err = global.DB.Take(&jt, id).Error
	if err != nil {
		res.FailWithMessage("跳转名称不存在", c)
		return
	}

	err = global.DB.Model(&jt).Updates(map[string]any{
		"jump_target_name": JTR.JumpTargetName,
		"jump_target_url":  JTR.JumpTargetURL,
		"images":           JTR.Images,
		"is_show":          JTR.IsShow,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改跳转链接失败", c)
		return
	}

	res.OKWithMessage("修改跳转链接成功", c)
}
