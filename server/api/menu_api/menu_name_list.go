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

func (MenuApi) MenuNameListView(c *gin.Context) {
	var menuNameList []NameListResponse

	err := global.DB.Model(models.MenuModels{}).Select("id", "menu_title", "path").Find(&menuNameList).Error
	if err != nil {
		res.FailWithError(err, menuNameList, c)
	}
	res.OKWithData(menuNameList, c)
	return
}
