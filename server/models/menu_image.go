package models

type MenuImages struct {
	MenuID     uint        `gorm:"primaryKey" json:"menu_id"`
	MenuModel  MenuModels  `gorm:"foreignKey:MenuID"`
	ImageID    uint        `gorm:"primaryKey" json:"image_id"`
	ImageModel ImageModels `gorm:"foreignKey:ImageID"`
	Sort       int         `json:"sort"`
}
