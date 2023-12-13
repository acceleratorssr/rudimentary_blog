package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

func (ImagesApi) ImageDeleteView(c *gin.Context) {
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
