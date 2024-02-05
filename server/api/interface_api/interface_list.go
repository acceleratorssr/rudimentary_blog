package interface_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/service/common"
)

// InterfaceListView 查询接口列表
//
// @Tags 接口
// @Summary  接口列表
// @Description 查询接口列表
// @Param data query models.Page false "查询参数"
// @Accept  json
// @Router /api/interface_list [get]
// @Produce json
// @Success 200 {object} res.Response
func (InterfaceApi) InterfaceListView(c *gin.Context) {
	var page models.Page
	var interfaceList []models.InterfaceModels
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	//var choose string
	//// admin:1 user:2 normal:3 banned:4
	//switch permission.Permissions {
	//case 1:
	//	choose = "ID|CreatedAt|UpdatedAt|Username|NickName|Password|Avatar|Token|IP|PhoneNum|Permission|SignStatus"
	//case 2:
	//	choose = "ID|CreatedAt|UpdatedAt|NickName|Avatar|IP|SignStatus"
	//case 3:
	//	choose = "ID|CreatedAt|UpdatedAt|NickName|Avatar"
	//default:
	//
	//}

	choose := "ID|CreatedAt|UpdatedAt|InterfaceName|UserId|Description|Url|Method|RequestHeader|ResponseHeader|Status"
	// 可继续进行脱敏操作
	totalPages, flag := common.ComList(models.InterfaceModels{}, page, &interfaceList, choose, c)
	if flag {
		res.OKWithList(interfaceList, totalPages, c)
	}

	return
}
