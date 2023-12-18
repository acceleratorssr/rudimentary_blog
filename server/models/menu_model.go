package models

import "server/models/status_type"

type MenuModels struct {
	MODEL
	MenuTitle   string        `gorm:"size:32" json:"menu_title"`
	MenuTitleEn string        `gorm:"size:32" json:"menu_title_en"`
	Path        string        `gorm:"size:64" json:"path"`
	MenuIcon    string        `gorm:"size:32" json:"menu_icon"`
	MenuImages  []ImageModels `gorm:"many2many:menu_images;joinForeignKey:MenuID;JoinReferences:ImageID" json:"menu_images"`
	// 菜单图片的切换间隔时间
	MenuTime int `json:"menu_time"`
	Abstract status_type.Array
	ParentId int `gorm:"size:11" json:"parent_id"`
	Sort     int `gorm:"size:11" json:"sort"`
}
