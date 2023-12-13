package settings_api

import (
	"github.com/gin-gonic/gin"
	"server/core"
	"server/global"
	"server/models/res"
)

// 单个:

//func (s SettingsApi) SettingsUpdate(c *gin.Context) {
//	var cs map[string]interface{}
//	// 将 HTTP 请求中的 JSON 数据绑定到一个 Go 的map中（结构体也可以）
//	err := c.ShouldBindJSON(&cs)
//	if err != nil {
//		res.FailWithCode(res.ParamsError, c)
//		return
//	}
//
//	// 定义需要更新的字段列表
//	// 此处对应的是json字段
//	fieldsToUpdate := []string{"created_at", "title", "email", "name", "addr", "githubUrl"}
//
//	fmt.Println("before", global.Config.SiteInfo)
//	// 遍历字段，仅更新非空字段
//	for _, field := range fieldsToUpdate {
//		if val, ok := cs[field]; ok && val != nil {
//			// 更新配置信息
//			switch field {
//			case "created_at":
//				if v, ok := val.(string); ok {
//					global.Config.SiteInfo.CreatedAt = v
//				}
//			case "title":
//				if v, ok := val.(string); ok {
//					global.Config.SiteInfo.Title = v
//				}
//			case "email":
//				if v, ok := val.(string); ok {
//					global.Config.SiteInfo.Email = v
//				}
//			case "name":
//				if v, ok := val.(string); ok {
//					global.Config.SiteInfo.Name = v
//				}
//			case "addr":
//				if v, ok := val.(string); ok {
//					global.Config.SiteInfo.Addr = v
//				}
//			case "githubUrl":
//				if v, ok := val.(string); ok {
//					global.Config.SiteInfo.GithubUrl = v
//				}
//			}
//		}
//	}
//	fmt.Println("after", global.Config.SiteInfo)
//	err = core.UpdateYaml()
//	if err != nil {
//		global.Log.Error(err)
//		res.FailWithMessage(err.Error(), c)
//		return
//	}
//	res.OK(c)
//}

type FieldUpdater struct {
	Field      string
	UpdateFunc func(interface{})
}

func (s SettingsApi) SettingsUpdate(c *gin.Context) {
	var cs map[string]string
	err := c.ShouldBindJSON(&cs)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	// 定义需要更新的字段列表及其更新函数
	fieldsToUpdate := []FieldUpdater{
		{"created_at", func(v interface{}) {
			if val, ok := v.(string); ok {
				global.Config.SiteInfo.CreatedAt = val
			}
		}},
		{"title", func(v interface{}) {
			if val, ok := v.(string); ok {
				global.Config.SiteInfo.Title = val
			}
		}},
		{"email", func(v interface{}) {
			if val, ok := v.(string); ok {
				global.Config.SiteInfo.Email = val
			}
		}},
		{"name", func(v interface{}) {
			if v, ok := v.(string); ok {
				global.Config.SiteInfo.Name = v
			}
		}},
		{"addr", func(v interface{}) {
			if v, ok := v.(string); ok {
				global.Config.SiteInfo.Addr = v
			}
		}},
		{"githubUrl", func(v interface{}) {
			if v, ok := v.(string); ok {
				global.Config.SiteInfo.GithubUrl = v
			}
		}},
		// 添加其他字段的更新函数
	}

	// 遍历字段，仅更新非空字段
	f := false
	for _, fieldUpdater := range fieldsToUpdate {
		if val, ok := cs[fieldUpdater.Field]; ok && val != "" {
			// 调用更新函数
			fieldUpdater.UpdateFunc(val)
			f = true
		}
	}
	if !f {
		res.FailWithMessage("not find this field -> SettingsUpdate", c)
	} else {
		err = core.UpdateYaml()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage(err.Error(), c)
			return
		}
		res.OK(c)
	}
}
