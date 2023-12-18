package models

import (
	"server/models/status_type"
)

type ArticleModels struct {
	MODEL
	Title string `gorm:"size:32;not null" json:"title"`
	// utf-8：一个中文字符占用3个字节
	Abstract string `gorm:"type=varchar(3000)" json:"abstract"`
	Content  string `gorm:"type=varchar(30000)" json:"content"`
	AuthorID uint   `json:"author_id"`
	// 这里设置外键为AuthorID
	Author   UserModels `gorm:"foreignKey:AuthorID" json:"author"`
	Category string     `gorm:"size:32:" json:"category"`

	PageView   int `gorm:"default:0" json:"page_view"`
	Like       int `gorm:"default:0" json:"like"`
	Comment    int `gorm:"default:0" json:"comment"`
	Collection int `gorm:"default:0" json:"collection"`

	IsTop       bool `gorm:"default:false" json:"is_top"`
	IsSecret    bool `gorm:"default:false" json:"is_secret"`
	IsHot       bool `gorm:"default:false" json:"is_hot"`
	IsRecommend bool `gorm:"default:false" json:"is_recommend"`

	CommentModels []CommentModels `gorm:"foreignKey:PostID" json:"-"`
	TagsModel     []TagsModels    `gorm:"many2many:article_tag" json:"tags_model"`

	Tags status_type.Array `gorm:"type:string;size:64" json:"tag"`

	// ArticleID uint   `json:"article_id"`
	//Photo   []ImageModels `gorm:"foreignKey:ArticleID" json:"-"`
	PhotoID uint `json:"PhotoID"`
}
