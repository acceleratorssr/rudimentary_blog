package jumpTarget_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/service/common"
	"strings"
)

func (JumpTargetApi) JumpTargetListView(c *gin.Context) {
	var jt models.Page
	var jumpTargetList []models.JumpTargetModel
	err := c.ShouldBindQuery(&jt)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	referer := c.GetHeader("Referer")
	isShow := true
	// 等于说有admin就不做筛选
	if strings.Contains(referer, "admin") {
		isShow = false
	}
	// true对应1
	// 可以利用表内字段做筛查
	// 注意：gorm特性：JumpTargetModel{IsShow: false}会被忽略，只能用true筛选
	totalPages, flag := common.ComList(models.JumpTargetModel{IsShow: isShow}, jt, &jumpTargetList, c)
	if flag {
		res.OKWithList(jumpTargetList, totalPages, c)
	}
}
