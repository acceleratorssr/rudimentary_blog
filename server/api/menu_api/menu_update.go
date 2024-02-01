package menu_api

import (
	myutils "github.com/acceleratorssr/My_go_utils"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/models/stype"
	"strings"
)

type son struct {
	MenuTitle   string      `json:"menu_title"`
	MenuTitleEn string      `json:"menu_title_en"`
	Path        string      `json:"path"`
	MenuIcon    string      `json:"menu_icon"`
	MenuTime    int         `json:"menu_time"`
	Abstract    stype.Array `json:"abstract"`
	ParentId    int         `json:"parent_id"`
	Sort        int         `json:"sort"`
	ImageSort   []ImageSort `json:"image_sort"`
}

type MenuUpdateRequest struct {
	Son          son    `json:"son"`
	FieldBanList string `json:"field_ban_list"`
}

// MenuUpdateView 更新菜单
//
// @Tags 菜单
// @Summary  更新菜单
// @Description 更新菜单的字段或者对应图片
// @Param id query string true "需要更新菜单序号"
// @Param data body MenuUpdateRequest true "需要更新的字段"
// @Accept  json
// @Router /api/menu [put]
// @Produce json
// @Success 200 {object} res.Response
func (MenuApi) MenuUpdateView(c *gin.Context) {
	var MR MenuUpdateRequest
	var menu models.MenuModels
	err := c.ShouldBindJSON(&MR)
	if err != nil {
		res.FailWithError(err, &MR, c)
		return
	}

	id := c.Param("id")
	if id == "" {
		res.FailWithMessage("菜单ID不能为空", c)
		return
	}

	err = global.DB.Take(&menu, id).Error
	if err != nil {
		res.FailWithMessage("菜单id不存在", c)
		return
	}

	//fieldBanList := "MenuTitleEnPathMenuIconMenuTimeAbstractParentIdSortImageSort"
	// 如果不更新图片 -> 不更新连接表
	if strings.Contains(MR.FieldBanList, "ImageSort") {
		err = global.DB.Model(&menu).Updates(myutils.StructToMap(MR.Son, MR.FieldBanList)).Error
	} else {
		err = global.DB.Model(&menu).Association("MenuImages").Clear()
		if err != nil {
			res.FailWithMessage("未找到对应菜单和图片的关系", c)
			return
		}

		if len(MR.Son.ImageSort) != 0 {
			var imageList []models.MenuImages
			for _, v := range MR.Son.ImageSort {
				imageList = append(imageList, models.MenuImages{
					MenuID:  menu.ID,
					ImageID: v.ImageID,
					Sort:    v.Sort,
				})
			}

			err = global.DB.Create(&imageList).Error
			if err != nil {
				global.Log.Error(err)
				res.FailWithMessage("更新图片失败", c)
				return
			}
		}

		err = global.DB.Model(&menu).Updates(myutils.StructToMap(MR.Son, MR.FieldBanList+"ImageSort")).Error
	}
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}

	res.OKWithMessage("修改菜单成功", c)
}
