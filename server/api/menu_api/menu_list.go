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
	MenuModels models.MenuModels
	Images     []Image `json:"images"`
}

// MenuList 菜单列表视图
//
// @Summary 获取菜单列表
// @Description 获取所有菜单的列表
// @Tags 菜单
// @Accept json
// @Produce json
// @Success 200 {array} MenuResponse "成功响应"
// @Router /api/menu [get]
func (MenuApi) MenuList(c *gin.Context) {
	var menuList []models.MenuModels
	var menuIDList []uint
	err := global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList).Error
	if err != nil {
		res.FailWithError(err, menuList, c)
		return
	}

	// 查连接表
	var menuImages []models.MenuImages
	// 连表查询Preload("ImageModel")
	global.DB.Preload("ImageModel").Order("sort desc").Find(&menuImages, "menu_id in (?)", menuIDList)
	var menus []MenuResponse

	// 根据menuID连表查询，返回对应的imageID，降序
	for _, v := range menuList {
		// 未初始化的切片默认为 nil
		//var images []Image
		// 这样写会被初始化为一个空切片([])，而不是 nil
		//images := []Image{}
		// 使用 make 函数创建切片时，会分配底层数组并返回一个包含指向该数组的切片
		// 即使长度为0，底层数组也会被分配
		images := make([]Image, 0)
		for _, vv := range menuImages {
			if vv.MenuID == v.ID {
				images = append(images, Image{
					ID:   vv.ImageID,
					Path: vv.ImageModel.Path,
				})
			}
		}
		menus = append(menus, MenuResponse{
			MenuModels: v,
			Images:     images,
		})
	}
	res.OKWithData(menus, c)
	return
}
