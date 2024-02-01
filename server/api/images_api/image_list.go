package images_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/service/common"
)

type ImageModelForSwagger struct {
	models.MODEL
	Path string `json:"path"`
	Key  string `json:"key"`
	Name string `gorm:"sizeof:32" json:"name"`
	// 不包含ImageMenus字段
}

// ImageListView
//
// @Summary 获取图片列表
// @Description 根据分页参数获取图片列表
// @Tags 图片
// @Accept json
// @Produce json
// @Param page query int true "页码"
// @Param limit query int true "每页数量"
// @Success 200 {array} models.ImageModels
// @Router /api/images [get]
func (ImagesApi) ImageListView(c *gin.Context) {
	// 请求方法：
	// http://127.0.0.1:9190/api/images?page=1&limit=2
	var imageList []models.ImageModels
	var p models.Page
	// 绑定查询参数到结构体
	err := c.ShouldBindQuery(&p)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	choose := "ID|CreatedAt|UpdatedAt|Path|Key|Name"
	// 传models时不需要传&，传imageList要加&，这样内部变化才影响外部
	totalPages, flag := common.ComList(models.ImageModels{}, p, &imageList, choose, c)

	if flag {
		res.OKWithList(imageList, totalPages, c)
	}
	return
}
