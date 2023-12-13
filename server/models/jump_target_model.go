package models

type JumpTargetModel struct {
	MODEL
	JumpTargetName string `json:"jump_target_name"`
	JumpTargetURL  string `json:"jump_target_url"`
	Images         string `json:"images"`
	IsShow         bool   `json:"is_show"`
}
