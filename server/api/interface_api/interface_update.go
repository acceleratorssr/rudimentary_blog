package interface_api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"server/global"
	"server/models"
	"server/models/res"
)

type InterfaceUpdateRequest struct {
	ID             uint   `json:"id"`
	InterfaceName  string `json:"interface_name"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Method         string `json:"method"`
	RequestHeader  string `json:"request_header"`
	ResponseHeader string `json:"response_header"`
	Status         string `json:"status"`
}

// 是否需要唯一，false表示不需要唯一
func isOnly(only []string, n string) bool {
	for _, s := range only {
		if s == n {
			return true
		}
	}
	return false
}

// InterfaceUpdate 是一个API视图，用于处理更新接口的请求
//
// @Summary 管理员更新接口信息
// @Description 接口更新视图
// @Tags 接口
// @Accept json
// @Produce json
// @Param InterfaceUpdateRequest body InterfaceUpdateRequest true "用户ID，可改接口名称、描述、url、请求方法、请求头、响应头、接口状态"
// @Success 200 {string} string "修改成功"
// @Router /api/interface_update [post]
func (InterfaceApi) InterfaceUpdate(c *gin.Context) {
	onlyOne := []string{"Url"}
	var IUR InterfaceUpdateRequest
	err := c.ShouldBindJSON(&IUR)
	if err != nil {
		global.Log.Warn("InterfaceUpdate -> 参数错误:", err)
		res.FailWithMessage("InterfaceUpdate -> 参数错误", c)
		return
	}

	var interfaces models.InterfaceModels
	err = global.DB.Take(&interfaces, "url = ?", IUR.Url).Error
	if err != nil {
		res.FailWithMessage("UserUpdate -> 接口不存在", c)
		return
	}

	iurType := reflect.TypeOf(IUR)
	iurVal := reflect.ValueOf(IUR)
	if iurType.Kind() == reflect.Ptr {
		iurVal = iurVal.Elem()
		iurType = iurType.Elem()
	}

	for i := 0; i < iurVal.NumField(); i++ {
		f := iurVal.Field(i)
		// 判断字段是否为空
		if !f.IsZero() {
			if isOnly(onlyOne, iurType.Field(i).Name) {
				var u models.InterfaceModels
				// 根据字段名查询数据库中models.InterfaceModels表中是否存在相同字段值的记录。
				global.DB.Where(iurType.Field(i).Name+" = ?", f.Interface()).First(&u)
				// 若存在且ID与传入的interfaces.ID不同且不为0，则输出日志并返回错误信息
				if u.ID != interfaces.ID && u.ID != 0 {
					global.Log.Warn("InterfaceUpdate -> url重复:", iurType.Field(i).Name)
					res.FailWithMessage("InterfaceUpdate -> url重复", c)
					return
				}
				// 若不存在冲突记录，则尝试更新数据库中的对应字段值
				err = global.DB.Model(&interfaces).Update(iurType.Field(i).Name, f.Interface()).Error
				if err != nil {
					global.Log.Warn("InterfaceUpdate -> 修改失败:", err)
					res.FailWithMessage("InterfaceUpdate -> 修改失败", c)
					return
				}
			} else {
				err = global.DB.Model(&interfaces).Update(iurType.Field(i).Name, f.Interface()).Error
				if err != nil {
					global.Log.Warn("InterfaceUpdate -> 修改失败:", err)
					res.FailWithMessage("InterfaceUpdate -> 修改失败", c)
					return
				}
			}
		}
	}

	res.OKWithAll(IUR, "UserUpdate -> 修改成功:", c)
}
