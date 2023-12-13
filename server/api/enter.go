package api

import (
	"server/api/images_api"
	"server/api/jumpTarget_api"
	"server/api/settings_api"
)

type Group struct {
	SettingsApi   settings_api.SettingsApi
	ImageApi      images_api.ImagesApi
	JumpTargetApi jumpTarget_api.JumpTargetApi
}

var Groups = new(Group)
