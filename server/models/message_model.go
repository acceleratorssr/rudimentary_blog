package models

import "time"

// MessageModels 貌似头像，昵称等消息暂时不用存储？
type MessageModels struct {
	CreatedAt time.Time `json:"created_at"`
	// 发送人ID
	SendUserID       uint       `gorm:"primaryKey" json:"send_user_id"`
	SendUserModel    UserModels `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName string     `gorm:"size:32" json:"send_user_nick_name"`
	SendUserAvatar   string     `gorm:"size:255" json:"send_user_avatar"`

	//接收者ID
	ReceiveUserID       uint       `gorm:"primaryKey" json:"receive_user_id"`
	ReceiveUserModel    UserModels `gorm:"foreignKey:ReceiveUserID" json:"-"`
	ReceiveUserNickName string     `gorm:"size:32" json:"receive_user_nick_name"`
	ReceiveUserAvatar   string     `gorm:"size:255" json:"receive_user_avatar"`

	// 消息内容
	Content string `gorm:"size:255" json:"content"`
}
