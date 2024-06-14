package menu_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

type NameListResponse struct {
	ID        uint   `json:"id"`
	MenuTitle string `json:"menu_title"`
	Path      string `json:"path"`
}

// MenuNameList 菜单名称列表视图
//
// @Summary 获取菜单名称列表
// @Description 获取所有菜单的名称列表
// @Tags 菜单
// @Accept json
// @Produce json
// @Success 200 {array} NameListResponse "成功响应"
// @Router /api/menuName [get]
func (MenuApi) MenuNameList(c *gin.Context) {
	var menuNameList []NameListResponse

	err := global.DB.Model(models.MenuModels{}).Select("id", "menu_title", "path").Find(&menuNameList).Error
	if err != nil {
		res.FailWithError(err, menuNameList, c)
	}
	res.OKWithData(menuNameList, c)
	return
}
