package menu_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/models/status_type"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	MenuTitle   string            `json:"menu_title" binding:"required" msg:"缺少菜单名称"`
	MenuTitleEn string            `json:"menu_title_en"`
	Path        string            `json:"path" binding:"required" msg:"缺少菜单路径"`
	MenuIcon    string            `json:"menu_icon"`
	MenuTime    int               `json:"menu_time"`
	Abstract    status_type.Array `json:"abstract"`
	ParentId    int               `json:"parent_id"`
	Sort        int               `json:"sort" binding:"required" msg:"缺少菜单排序"`
	ImageSort   []ImageSort       `json:"image_sort"`
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var MR MenuRequest
	err := c.ShouldBindJSON(&MR)
	if err != nil {
		res.FailWithError(err, &MR, c)
		return
	}

	var menu models.MenuModels
	err = global.DB.Select("menu_title").Take(&menu, "menu_title = ? or path = ?", MR.MenuTitle, MR.Path).Error
	if err == nil {
		res.FailWithMessage("菜单名称已存在", c)
		return
	}
	// 创建菜单
	menu = models.MenuModels{
		MenuTitle:   MR.MenuTitle,
		MenuTitleEn: MR.MenuTitleEn,
		Path:        MR.Path,
		MenuIcon:    MR.MenuIcon,
		MenuTime:    MR.MenuTime,
		Abstract:    MR.Abstract,
		ParentId:    MR.ParentId,
		Sort:        MR.Sort,
	}
	err = global.DB.Create(&menu).Error
	if err != nil {
		global.Log.Error("菜单创建失败", err)
		res.FailWithMessage("菜单创建失败", c)
		return
	}

	if len(MR.ImageSort) == 0 {
		res.OKWithMessage("菜单创建成功", c)
		return
	}

	// 添加到关联表
	var menuImage []models.MenuImages
	// 不要这样写，:= 是用于声明并初始化变量的短变量声明语句
	//menuImage := []models.MenuImages

	for _, v := range MR.ImageSort {
		// 需要判断imageId是否存在
		menuImage = append(menuImage, models.MenuImages{
			MenuID:  menu.ID,
			ImageID: v.ImageID,
			Sort:    v.Sort,
		})
	}

	err = global.DB.Create(&menuImage).Error
	if err != nil {
		global.Log.Error("菜单图片关联失败", err)
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OKWithMessage("菜单创建成功", c)
}
