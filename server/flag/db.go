package flag

import (
	"fmt"
	"server/global"
	"server/models"
)

func Makemigrations() {
	err := global.DB.SetupJoinTable(&models.UserModels{}, "CollectsModels", &models.UserCollections{})
	if err != nil {

		return
	}
	// 注意在models.MenuModel表中，必须对应有第二个参数："MenuImages"字段（或者其他名字），并且设置好多对多 等的选项
	// 注意设置的外键名字和对应连接表的外键名字要相同
	err = global.DB.SetupJoinTable(&models.MenuModels{}, "MenuImages", &models.MenuImages{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 生成表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.ArticleModels{},
			&models.UserModels{},
			&models.FeedbackModels{},
			&models.CommentModels{},
			&models.ImageModels{},
			&models.MessageModels{},
			&models.TagsModels{},
			&models.JumpTargetModels{},
			&models.MenuModels{},
			&models.UserCollections{},
			&models.ArticleModels{},
			&models.LoginDataModels{},
			&models.InterfaceModels{},
		)
	if err != nil {
		global.Log.Errorf("Makemigrations fail:%s", err)
		return
	}
	global.Log.Info("Makemigrations success")
}
