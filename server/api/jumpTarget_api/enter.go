package jumpTarget_api

import "server/models"

type JumpTargetApi struct {
	models.MODEL
	Title    string `json:"title"`
	ImageUrl string `json:"image_url"`
	Url      string `json:"url"`
}

type TargetList struct {
}
