package models

type FeedbackModels struct {
	MODEL
	Email        string `gorm:"size:64" json:"email"`
	Content      string `gorm:"size:300" json:"content"`
	ApplyContent string `gorm:"size:300" json:"apply_content"`
	IsApply      bool   `json:"is_apply"`
}
