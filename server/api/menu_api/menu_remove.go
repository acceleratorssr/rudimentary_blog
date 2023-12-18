package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/global"
	"server/models"
	"server/models/res"
)

// MenuRemoveView 删除菜单
// @Tags 菜单
// @Summary  删除菜单
// @Description 删除多或者单个菜单
// @Param data body models.RemoveRequest true "需要删除的id_list"
// @Accept  json
// @Router /api/menu [delete]
// @Produce json
// @Success 200 {object} res.Response
func (MenuApi) MenuRemoveView(c *gin.Context) {
	var RQ models.RemoveRequest
	err := c.ShouldBindJSON(&RQ)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var menuList []models.MenuModels
	count := global.DB.Find(&menuList, RQ.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("对应菜单不存在", c)
		return
	}

	// 开启事务，自动回滚
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除连接表内数据
		err = global.DB.Model(&menuList).Association("MenuImages").Clear()
		if err != nil {
			res.OKWithMessage("对应菜单和图片联系不存在", c)
			return err
		}
		// 数据库删除
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			res.FailWithMessage("删除菜单失败", c)
			return err
		}
		return err
	})
	// 如果进入if，则事务没有执行
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}

	res.OKWithMessage(fmt.Sprintf("成功删除%d个菜单", count), c)
}
