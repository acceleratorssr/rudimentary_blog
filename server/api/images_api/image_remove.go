package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// ImageRemoveView 删除文件
// @Tags 图片
// @Summary  删除对应的文件
// @Description 例如：image
// @Param data body models.RemoveRequest true "需要删除的id_list"
// @Accept  json
// @Router /api/images [delete]
// @Produce json
// @Success 200 {object} res.Response
func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var RQ models.RemoveRequest
	err := c.ShouldBindJSON(&RQ)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var imageList []models.ImageModel
	count := global.DB.Find(&imageList, RQ.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	// 数据库删除
	global.DB.Delete(&imageList)
	res.OKWithMessage(fmt.Sprintf("成功删除%d份图片", count), c)
}
