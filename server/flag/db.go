package flag

import (
	"fmt"
	"server/global"
	"server/models"
)

func Makemigrations() {
	err := global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollection{})
	if err != nil {

		return
	}
	// 注意在models.MenuModel表中，必须对应有第二个参数："MenuImage"字段（或者其他名字），并且设置好多对多 等的选项
	// 注意设置的外键名字和对应连接表的外键名字要相同
	err = global.DB.SetupJoinTable(&models.MenuModel{}, "MenuImage", &models.MenuImage{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.ArticleModel{},
			&models.UserModel{},
			&models.FeedbackModel{},
			&models.CommentModel{},
			&models.ImageModel{},
			&models.MessageModel{},
			&models.TagsModel{},
			&models.JumpTargetModel{},
			&models.MenuModel{},
			&models.UserCollection{},
			&models.ArticleModel{},
			&models.MenuImage{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Log.Errorf("Makemigrations fail:%s", err)
		return
	}
	global.Log.Info("Makemigrations success")
}
