package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// ImageUpdateView 更新文件名
// @Tags 图片
// @Summary  更新对应的文件名
// @Description 例如：image
// @Param data body models.UpdateRequest true "需要更新的id以及对应新的名称"
// @Accept  json
// @Router /api/images [put]
// @Produce json
// @Success 200 {object} res.Response
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var UR models.UpdateRequest
	// UR设置了binding:"required"和对应报错msg
	// res.FailWithError(err, UR, c)会输出字段中自定义msg
	err := c.ShouldBindJSON(&UR)
	if err != nil {
		res.FailWithError(err, UR, c)
		return
	}

	var imageModel models.ImageModel
	err = global.DB.Take(&imageModel, UR.ID).Error
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("文件不存在 :%s", err), c)
		return
	}

	err = global.DB.Model(&imageModel).Update("name", UR.Name).Error
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("name :%s", err), c)
		return
	}
	res.OKWithMessage("文件名修改成功", c)
	return
}
