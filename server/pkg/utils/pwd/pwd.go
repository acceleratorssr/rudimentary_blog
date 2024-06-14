package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"server/global"
)

// HashAndSalt 加密密码
func HashAndSalt(pwd string) string {
	bytePwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		global.Log.Error("加密错误")
		return ""
	}
	return string(hash)
}

// CheckPasswords 验证密码
func CheckPasswords(hashedPwd string, rePwd string) bool {
	byteHash := []byte(hashedPwd)
	byteRePwd := []byte(rePwd)

	err := bcrypt.CompareHashAndPassword(byteHash, byteRePwd)
	if err != nil {
		global.Log.Warn("密码错误")
		return false
	}
	return true
}
