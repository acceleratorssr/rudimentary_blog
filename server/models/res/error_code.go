package res

type ErrorCode int

// 定义错误码和错误消息映射

const (
	SettingsError ErrorCode = 600 + iota
	ParamsError
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "settingsError",
		ParamsError:   "ParamsError",
	}
)
