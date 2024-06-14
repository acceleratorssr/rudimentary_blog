package images_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

type NameListResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// ImageNameList
//
// @Summary 获取图片名字列表
// @Description 返回所有图片名字
// @Tags 图片
// @Produce json
// @Success 200 {array} NameListResponse
// @Router /api/imagesName [get]
func (ImagesApi) ImageNameList(c *gin.Context) {
	var imageNameList []NameListResponse

	err := global.DB.Model(models.ImageModels{}).Select("id", "name", "path").Find(&imageNameList).Error
	if err != nil {
		res.FailWithError(err, imageNameList, c)
	}
	res.OKWithData(imageNameList, c)
	return
}
