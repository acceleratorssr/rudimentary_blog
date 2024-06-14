package jumpTarget_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

type JumpTargetRequest struct {
	JumpTargetName string `json:"jump_target_name" binding:"required" msg:"请输入跳转名称"`
	JumpTargetURL  string `json:"jump_target_url" binding:"required,url" msg:"非法跳转路径"`
	Images         string `json:"images" binding:"required" msg:"请输入所示图片"`
	IsShow         bool   `json:"is_show"`
}

// JumpTargetCreate 添加跳转的目标
//
// @Tags 跳转的目标
// @Summary  添加跳转目标
// @Description 添加跳转的目标
// @Param data body JumpTargetRequest true "表示多个参数"
// @Accept  json
// @Router /api/jumpTarget [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (JumpTargetApi) JumpTargetCreate(c *gin.Context) {
	var jtr JumpTargetRequest
	err := c.ShouldBindJSON(&jtr)
	if err != nil {
		res.FailWithError(err, jtr, c)
		return
	}

	var jt models.JumpTargetModels
	err = global.DB.Take(&jt, "jump_target_name = ?", jtr.JumpTargetName).Error
	if err == nil {
		res.FailWithMessage("跳转名称已存在", c)
		return
	}
	err = global.DB.Create(&models.JumpTargetModels{
		JumpTargetName: jtr.JumpTargetName,
		JumpTargetURL:  jtr.JumpTargetURL,
		Images:         jtr.Images,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("创建跳转链接失败", c)
		return
	}

	res.OKWithMessage("创建跳转链接成功", c)
}
