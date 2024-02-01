package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// MenuDetailView 菜单详情视图
//
// @Summary 获取菜单详情
// @Description 通过ID获取菜单详情
// @Tags 菜单
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Success 200 {object} MenuResponse "成功响应"
// @Router /api/menuDetail/{id} [get]
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
