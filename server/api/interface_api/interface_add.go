package interface_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/pkg/utils/jwts"
)

type InterfaceAddRequest struct {
	InterfaceName  string `json:"interface_name" binding:"required" msg:"缺少接口名"`
	UserId         uint   `json:"user_id"`
	Description    string `json:"description" `
	Url            string `json:"url" binding:"required,url" msg:"接口url有误"`
	Method         string `json:"method" binding:"required" msg:"缺少接口方法"`
	RequestHeader  string `json:"request_header"`
	ResponseHeader string `json:"response_header"`
	Status         string `json:"status"`
}

// InterfaceAdd 添加接口
//
// @Tags 接口
// @Summary  添加接口
// @Description 添加接口，url唯一
// @Param InterfaceAddRequest body InterfaceAddRequest true "创建接口参数"
// @Accept  json
// @Router /api/interface_add [post]
// @Produce json
// @Success 200 {object} models.InterfaceModels
func (InterfaceApi) InterfaceAdd(c *gin.Context) {
	var IAR InterfaceAddRequest
	var interfaceModel models.InterfaceModels

	err := c.ShouldBindJSON(&IAR)
	if err != nil {
		global.Log.Warnln("创建接口失败 InterfaceAdd -> ", err)
		res.FailWithError(err, InterfaceAddRequest{}, c)
		return
	}

	_permission, _ := c.Get("parseToken")

	// 注意_permission的类型是 *jwts.Permission
	permission := _permission.(*jwts.CustomClaims)

	// 注册接口到mysql，获取id
	err = global.DB.Take(&interfaceModel, "url = ?", IAR.Url).Error
	if err == nil {
		global.Log.Warn("接口注册 -> 接口已存在", err)
		res.FailWithMessage("接口已存在", c)
		return
	}

	interfaceApi := models.InterfaceModels{
		UserId:         permission.UserID,
		InterfaceName:  IAR.InterfaceName,
		Description:    IAR.Description,
		Url:            IAR.Url,
		Method:         IAR.Method,
		RequestHeader:  IAR.RequestHeader,
		ResponseHeader: IAR.ResponseHeader,
		Status:         IAR.Status,
	}

	err = global.DB.Create(&interfaceApi).Error
	if err != nil {
		global.Log.Warn("接口注册 -> 接口注册失败", err)
		res.FailWithMessage("接口注册失败", c)
		return
	}

	var result models.InterfaceModels
	err = global.DB.First(&result, "url = ?", interfaceApi.Url).Error
	if err != nil {
		global.Log.Warn("接口注册 -> 接口注册入数据库后返回数据失败", err)
		res.FailWithMessage("接口注册成功，返回数据失败", c)
		return
	}

	res.OKWithData(result, c)
}
