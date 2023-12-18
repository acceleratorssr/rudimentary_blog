package models

import (
	"gorm.io/gorm"
	"os"
	"server/global"
)

type ImageModels struct {
	MODEL
	Path       string       `json:"path"`
	Key        string       `json:"key"`
	Name       string       `gorm:"sizeof:32" json:"name"`
	ImageMenus []MenuModels `gorm:"many2many:menu_images;joinForeignKey:ImageID;JoinReferences:MenuID" json:"image_menus"`
	//ArticleID uint   `json:"article_id"`
	//ArticleModels []ArticleModels `gorm:"foreignKey:ArticleID" json:"-"`
	//UserModels    []UserModels
}

// BeforeDelete 刚刚好，这里的逻辑是多个不同的用户可以上传图片到数据库，并且都可以将图片从数据库中删除
// 如果本地存在和数据库对应的需要删除的图片时，也一起删除
// 如果本地没有对应的图片，那么数据库内的图片也不会被删除
// 简单来说：只有自己上传的图片才能被删除（目前没有开放没有下载他人图片）
// 如果返回一个非空的错误，GORM 将停止删除操作
func (i *ImageModels) BeforeDelete(tx *gorm.DB) (err error) {
	//filePathWithName := path.Join(i.Path, i.Name)
	err = os.Remove(i.Path)
	if err != nil {
		global.Log.Error(err)
		return err
	}

	return nil
}

//func (i *ImageModels) AfterUpdate(tx *gorm.DB) (err error) {
//
//	return nil
//}
