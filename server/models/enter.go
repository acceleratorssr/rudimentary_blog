package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Page form 指定了对应字段对应的查询参数名
// 这样无论大小写敏感与否，都能正确地绑定到结构体字段。
type Page struct {
	Page  int    `form:"page"`
	Key   int    `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type UpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件新名称"`
}
