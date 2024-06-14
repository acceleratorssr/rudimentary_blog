package jwts

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"server/global"
	"time"
)

type JwtPayload struct {
	Username    string `json:"username"`
	UserID      uint   `json:"user_id"`
	NickName    string `json:"nick_name"`
	Permissions int    `json:"permissions"`
}

type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims // 实现Claims接口的struct
}

var MySecret []byte

// GenToken 创建Token
func GenToken(user JwtPayload) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	//claim := jwt.MapClaims{
	//	"username":    user.Username,
	//	"user_id":     user.UserID,
	//	"nick_name":   user.NickName,
	//	"permissions": user.Permissions,
	//}
	////or
	claim := CustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
