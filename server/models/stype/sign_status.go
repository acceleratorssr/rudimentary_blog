package stype

import "encoding/json"

type SignStatus int

const (
	SignNotStatus = 1000 + iota
	SignPhoneNum
	SignWechat
	SignGithub
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	switch s {
	case SignNotStatus:
		return "NotStatusUP"
	case SignPhoneNum:
		return "PhoneNumUP"
	case SignWechat:
		return "WechatUP"
	case SignGithub:
		return "GithubUP"
	default:
		return "Unknown"
	}
}
