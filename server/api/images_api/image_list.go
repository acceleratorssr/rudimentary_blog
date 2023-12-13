package images_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/service/common"
)

func (ImagesApi) ImageListView(c *gin.Context) {
	// 请求方法：
	// http://127.0.0.1:9190/api/images?page=1&limit=2
	var imageList []models.ImageModel
	var p models.Page
	// 绑定查询参数到结构体
	err := c.ShouldBindQuery(&p)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	// 传models时不需要传&，传imageList要加&，这样内部变化才影响外部
	totalPages, flag := common.ComList(models.ImageModel{}, p, &imageList, c)

	if flag {
		res.OKWithList(imageList, totalPages, c)
	}
	return
}
