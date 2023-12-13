package utils

func InList(key string, list []string) bool {
	// 检查文件后缀是否在允许的列表中
	for _, ext := range list {
		if key == ext {
			return true
		}
	}
	return false
}
