package models

import "time"

type UserCollections struct {
	UserID       uint          `gorm:"primaryKey"`
	UserModel    UserModels    `gorm:"foreignKey:UserID"`
	ArticleID    uint          `gorm:"primaryKey"`
	ArticleModel ArticleModels `gorm:"foreignKey:ArticleID"`
	CreateAt     time.Time
}
