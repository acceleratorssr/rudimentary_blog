package api

import (
	"server/api/images_api"
	"server/api/jumpTarget_api"
	"server/api/menu_api"
	"server/api/message_api"
	"server/api/settings_api"
	"server/api/user_api"
)

type Group struct {
	SettingsApi   settings_api.SettingsApi
	ImageApi      images_api.ImagesApi
	JumpTargetApi jumpTarget_api.JumpTargetApi
	MenuApi       menu_api.MenuApi
	UserApi       user_api.UserApi
	MessageApi    message_api.MessageApi
}

var Groups = new(Group)
