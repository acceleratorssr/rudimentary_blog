package models

type TagsModels struct {
	MODEL
	Title        string          `gorm:"size:16" json:"title"`
	ArticleModel []ArticleModels `gorm:"many2many:article_tag" json:"-"`
}
