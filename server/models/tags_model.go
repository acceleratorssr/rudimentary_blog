package models

type TagsModel struct {
	MODEL
	Title        string         `gorm:"size:16" json:"title"`
	ArticleModel []ArticleModel `gorm:"many2many:article_tag" json:"-"`
}
