package models

import "server/models/status_type"

type MenuModel struct {
	MODEL
	MenuTitle string       `gorm:"size:32" json:"menu_title"`
	MenuIcon  string       `gorm:"size:32" json:"menu_icon"`
	MenuImage []ImageModel `gorm:"many2many:menu_image;joinForeignKey:MenuID;JoinReferences:ImageID" json:"menu_image"`
	// 菜单图片的切换间隔时间
	MenuTime int `json:"menu_time"`
	Abstract status_type.Array
	ParentId int `gorm:"size:11" json:"parent_id"`
	Sort     int `gorm:"size:11" json:"sort"`
}
