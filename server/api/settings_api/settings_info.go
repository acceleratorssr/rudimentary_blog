package settings_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/res"
)

type FieldInfo struct {
	Field      string
	UpdateFunc func()
}

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfo 是一个API视图，用于处理设置信息的请求
//
// @Summary 获取配置信息
// @Description 设置信息视图，需要具体配置名称。如果找不到该字段，将返回错误信息。
// @Tags 配置
// @Accept json
// @Produce json
// @Param name path string true "获取配置名称"
// @Success 200 {string} string "success"
// @Router /settings/{name} [get]
func (s SettingsApi) SettingsInfo(c *gin.Context) {
	//var cs string
	//// 获取 URI 参数 ":name"
	//cs = c.Param("name")
	//fmt.Println(cs)
	//或者：
	var cs SettingsUri
	// 记得加&
	// ShouldBindUri方法需要一个指向目标结构的指针，以便能够将 URI 参数的值绑定到结构体的字段上
	err := c.ShouldBindUri(&cs)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	// 注意大小写！
	fieldsInfo := []FieldInfo{
		{"mysql", func() {
			res.OKWithAll(global.Config.Mysql, "success", c)
		}},
		{"logger", func() {
			res.OKWithAll(global.Config.Logger, "success", c)
		}},
		{"system", func() {
			res.OKWithAll(global.Config.System, "success", c)
		}},
		{"site_info", func() {
			res.OKWithAll(global.Config.SiteInfo, "success", c)
		}},
		{"wechat", func() {
			res.OKWithAll(global.Config.Wechat, "success", c)
		}},
		{"qi_niu", func() {
			res.OKWithAll(global.Config.QiNiu, "success", c)
		}},
		{"jwt", func() {
			res.OKWithAll(global.Config.Jwt, "success", c)
		}},
		{"email", func() {
			res.OKWithAll(global.Config.Email163, "success", c)
		}},
	}

	f := false
	for _, fieldInformation := range fieldsInfo {
		//if fieldInformation.Field == cs {
		//	fieldInformation.UpdateFunc()
		//	f = true
		//}
		// 对应uri方法
		if fieldInformation.Field == cs.Name {
			fieldInformation.UpdateFunc()
			f = true
		}
	}
	if !f {
		res.FailWithMessage("not find this field -> SettingsInfo", c)
	} else {
		res.OK(c)
	}
}
