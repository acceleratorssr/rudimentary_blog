package models

type InterfaceModels struct {
	MODEL
	InterfaceName  string `gorm:"size:512;not null;unique" json:"interface_name"` //接口名字
	UserId         uint   `gorm:"not null" json:"user_id"`                        //创建人
	Description    string `gorm:"size:1024" json:"description"`
	Url            string `gorm:"size:512;not null;unique" json:"url"`
	Method         string `gorm:"not null" json:"method"` //请求类型
	RequestHeader  string `json:"request_header"`
	ResponseHeader string `json:"response_header"`
	Status         uint   `gorm:"default:0" json:"status"` //接口状态，0上线，1下线
}
