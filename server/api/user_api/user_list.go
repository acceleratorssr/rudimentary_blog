package user_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/models/stype"
	"server/service/common"
	"server/utils/jwts"
)

type userResponse struct {
	ID         uint             `json:"id"`
	NickName   string           `json:"nick_name"`
	Avatar     string           `json:"-"`
	IP         string           `json:"ip"`
	SignStatus stype.SignStatus `gorm:"type=smallint(6);not null" json:"sign_status"`
	//ArticleModels []ArticleModels        `gorm:"foreignKey:AuthorID" json:"-"`
	//CollectsModels []ArticleModels `gorm:"many2many:user_collection;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"`
}

// UserListView 查询用户列表
//
// @Tags 用户
// @Summary  用户列表
// @Description 查询用户列表
// @Param data query models.Page false "查询参数"
// @Accept  json
// @Router /api/user_list [get]
// @Produce json
// @Success 200 {object} res.Response
func (UserApi) UserListView(c *gin.Context) {
	_permission, _ := c.Get("parseToken")

	// 注意_permission的类型是 *jwts.Permission
	permission := _permission.(*jwts.CustomClaims)

	var page models.Page
	var userList []models.UserModels
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var choose string
	// admin:1 user:2 normal:3 banned:4
	switch permission.Permissions {
	case 1:
		choose = "ID|CreatedAt|UpdatedAt|Username|NickName|Password|Avatar|Token|IP|PhoneNum|Permission|SignStatus"
	case 2:
		choose = "ID|CreatedAt|UpdatedAt|NickName|Avatar|IP|SignStatus"
	case 3:
		choose = "ID|CreatedAt|UpdatedAt|NickName|Avatar"
	default:

	}
	// 可继续进行脱敏操作
	totalPages, flag := common.ComList(models.UserModels{}, page, &userList, choose, c)
	if flag {
		res.OKWithList(userList, totalPages, c)
	}

	return
}
