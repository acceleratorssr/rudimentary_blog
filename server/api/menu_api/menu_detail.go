package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	var menu models.MenuModels
	err := global.DB.Take(&menu, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	var images = make([]Image, 0)
	// 查连接表
	var menuImages []models.MenuImages
	global.DB.Preload("ImageModel").Order("sort desc").Find(&menuImages, "menu_id = ?", id)
	for _, v := range menuImages {
		if menu.ID == v.MenuID {
			images = append(images, Image{
				ID:   v.MenuID,
				Path: v.ImageModel.Path,
			})
		}
	}

	menusResponse := MenuResponse{
		MenuModels: menu,
		Images:     images,
	}

	res.OKWithData(menusResponse, c)
	return
}
