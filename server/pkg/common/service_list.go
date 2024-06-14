package common

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/global"
	"server/models"
	"server/models/res"
	"strings"
)

// ComList 通用的，根据表，page和limit进行分页
// 注意，list的类型是*[]T，不是[]T，如果是后者
// 在find之后，list会被置空，导致后续的查询失败
// choose为显示的字段
func ComList[T any](model T, page models.Page, list *[]T, choose string, c *gin.Context) (totalPages int, f bool) {
	//// 此处可开启日志
	DB := global.DB.Session(&gorm.Session{PrepareStmt: true, Logger: global.MysqlLog})

	chooseArr := strings.Split(choose, "|")

	//DB := global.DB
	//var totalCount int64
	//global.DB.Model(&models.ImageModels{}).Count(&totalCount)
	//等效：
	// Where(model)只能筛选true
	totalCount := int(DB.Where(model).Select("id").Find(&list).RowsAffected)
	//totalCount := int(DB.Where("is_show = ?", false).Find(&list).RowsAffected)
	if len(*list) == 0 {
		res.FailWithMessage("暂无符合的数据", c)
		return 0, false
	}

	if page.Limit == 0 || page.Page == 0 {
		DB.Find(&list)
		return 1, true
	}

	// 计算总页数，或者说有效页数
	totalPages = (totalCount + page.Limit - 1) / page.Limit

	// 如果page或者limit没有被初始化，则默认返回全部数据
	if totalPages < page.Page || page.Page < 0 {
		res.FailWithMessage("无效的页数", c)
		return 0, false
	}

	if page.Sort == "" {
		page.Sort = "created_at desc"
	}

	// 执行查询，限制每页记录数并设置偏移量
	DB.Select(chooseArr).Limit(page.Limit).Offset((page.Page - 1) * page.Limit).Order(page.Sort).Find(&list)
	// 返回总页数
	return totalPages, true
}
