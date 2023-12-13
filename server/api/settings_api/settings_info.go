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

func (s SettingsApi) SettingsInfoView(c *gin.Context) {
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
			res.OKWithAll(global.Config.Email, "success", c)
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
		res.FailWithMessage("not find this field -> SettingsInfoView", c)
	} else {
		res.OK(c)
	}
}
