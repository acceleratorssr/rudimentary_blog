package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

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
