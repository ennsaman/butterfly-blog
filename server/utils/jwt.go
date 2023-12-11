package utils

import (
	"blog-server/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MyClaims 定义 JWT 中保存的信息
type MyClaims struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type MyJWT struct {
	SecretKey []byte
}

func GetJWT() *MyJWT {
	return &MyJWT{[]byte(config.Conf.JWT.SecretKey)}
}

// GenToken 生成 JWT
func (j *MyJWT) GenToken(userId int, role string) (string, error) {
	// 创建一个 MyClaims 实例
	claims := MyClaims{
		UserId: userId,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.Conf.JWT.Issuer,
			ExpiresAt: time.Now().Add(time.Duration(config.Conf.JWT.ExpireTime)).Unix(),
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	// 使用指定的 secret 签名并获得完整的编码后的字符串 token
	return token.SignedString(j.SecretKey)
}
