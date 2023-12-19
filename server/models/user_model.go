package models

import (
	"server/models/status_type"
)

type UserModels struct {
	MODEL
	Username string `gorm:"size:32;not null" json:"user_name"`
	NickName string `gorm:"size=32" json:"nick_name"`
	Password string `gorm:"size:64;not null" json:"password"`
	// 可加默认头像
	Avatar   string `gorm:"size:256" json:"-"`
	Token    string `gorm:"64" json:"token"`
	IP       string `gorm:"size:20" json:"ip"`
	PhoneNum string `gorm:"size:11" json:"phone_num"`
	// admin:1 user:2 normal:3 banned:4
	Permission status_type.Permission `gorm:"size:4;not null;default:1" json:"permission"`

	SignStatus    status_type.SignStatus `gorm:"type=smallint(6);not null" json:"sign_status"`
	ArticleModels []ArticleModels        `gorm:"foreignKey:AuthorID" json:"-"`
	// joinForeignKey:UserID;JoinReferences:ArticleID这里名字要和连接表对应相同
	CollectsModels []ArticleModels `gorm:"many2many:user_collection;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"`
}
