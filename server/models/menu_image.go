package models

type MenuImage struct {
	MenuID     uint       `gorm:"primaryKey" json:"menu_id"`
	MenuModel  MenuModel  `gorm:"foreignKey:MenuID"`
	ImageID    uint       `gorm:"primaryKey" json:"image_id"`
	ImageModel ImageModel `gorm:"foreignKey:ImageID"`
	Sort       int        `json:"sort"`
}
