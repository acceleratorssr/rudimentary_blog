package status_type

import "encoding/json"

type SignStatus int

const (
	SignWechat = 1000 + iota
	SignGithub
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	switch s {
	case SignWechat:
		return "WechatUP"
	case SignGithub:
		return "GithubUP"
	default:
		return "Unknown"
	}
}
