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

func (JumpTargetApi) JumpTargetCreateView(c *gin.Context) {
	var JTR JumpTargetRequest
	err := c.ShouldBindJSON(&JTR)
	if err != nil {
		res.FailWithError(err, JTR, c)
		return
	}

	var jt models.JumpTargetModel
	err = global.DB.Take(&jt, "jump_target_name = ?", JTR.JumpTargetName).Error
	if err == nil {
		res.FailWithMessage("跳转名称已存在", c)
		return
	}
	err = global.DB.Create(&models.JumpTargetModel{
		JumpTargetName: JTR.JumpTargetName,
		JumpTargetURL:  JTR.JumpTargetURL,
		Images:         JTR.Images,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("创建跳转链接失败", c)
		return
	}

	res.OKWithMessage("创建跳转链接成功", c)
}
