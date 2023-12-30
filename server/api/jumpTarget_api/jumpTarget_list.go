package jumpTarget_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/service/common"
	"strings"
)

// JumpTargetListView 查询跳转的目标
// @Tags 跳转的目标
// @Summary  查询跳转目标
// @Description 查询跳转的目标
// @Param data query models.Page false "查询参数"
// @Accept  json
// @Router /api/jumpTarget [get]
// @Produce json
// @Success 200 {object} res.Response
func (JumpTargetApi) JumpTargetListView(c *gin.Context) {
	var jt models.Page
	var jumpTargetList []models.JumpTargetModels
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

	choose := "ID|CreatedAt|UpdatedAt|JumpTargetName|JumpTargetURL|Images|IsShow"

	// true对应1
	// 可以利用表内字段做筛查
	// 注意：gorm特性：JumpTargetModels{IsShow: false}会被忽略，只能用true筛选
	totalPages, flag := common.ComList(models.JumpTargetModels{IsShow: isShow}, jt, &jumpTargetList, choose, c)
	if flag {
		res.OKWithList(jumpTargetList, totalPages, c)
	}
}
