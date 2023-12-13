package models

type LoginDataModel struct {
	MODEL
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"20" json:"ip"`
	NickName  string    `gorm:"32" json:"nick_name"`
	Token     string    `json:"token"`
	Device    string    `gorm:"size:256" json:"device"`
	Addr      string    `gorm:"size:64" json:"addr"`
}
