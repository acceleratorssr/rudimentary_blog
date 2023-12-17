package menu_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

type Image struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Images []Image `json:"images"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	var menuList []models.MenuModel
	var menuIDList []uint
	err := global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList).Error
	if err != nil {
		res.FailWithError(err, menuList, c)
		return
	}

	// 查连接表
	var menuImages []models.MenuImage
	// 连表查询Preload("ImageModel")
	global.DB.Preload("ImageModel").Order("sort desc").Find(&menuImages, "menu_id in (?)", menuIDList)
	var menus []MenuResponse

	// 根据menuID连表查询，返回对应的imageID，降序
	for _, v := range menuList {
		var images []Image
		for _, vv := range menuImages {
			if vv.MenuID == v.ID {
				images = append(images, Image{
					ID:   vv.ImageID,
					Path: vv.ImageModel.Path,
				})
			}
		}
		menus = append(menus, MenuResponse{
			MenuModel: v,
			Images:    images,
		})
	}
	res.OKWithData(menus, c)
	return
}
