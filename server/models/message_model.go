package models

type MessageModel struct {
	MODEL
	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"`
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName string    `gorm:"size:32" json:"send_user_nick_name"`
	SendUserAvatar   string    `gorm:"size:255" json:"send_user_avatar"`

	ReceiveUserID       uint      `gorm:"primaryKey" json:"receive_user_id"`
	ReceiveUserModel    UserModel `gorm:"foreignKey:ReceiveUserID" json:"-"`
	ReceiveUserNickName string    `gorm:"size:32" json:"receive_user_nick_name"`
	ReceiveUserAvatar   string    `gorm:"size:255" json:"receive_user_avatar"`
	Content             string    `gorm:"size:255" json:"content"`
}
