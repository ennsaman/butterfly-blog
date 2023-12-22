package utils

import (
	"blog-server/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Token 错误
var (
	TokenExpired     = errors.New("token 已过期, 请重新登录")
	TokenNotValidYet = errors.New("token 尚未生效, 请重新登录")
	TokenMalformed   = errors.New("token 格式错误, 请重新登录")
	TokenInvalid     = errors.New("token 无效, 请重新登录")
)

// MyClaims 定义 JWT 中保存的信息
type MyClaims struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	UUID   string `json:"uuid"`
	jwt.StandardClaims
}

type MyJWT struct {
	SecretKey []byte
}

func GetJWT() *MyJWT {
	return &MyJWT{[]byte(config.Conf.JWT.SecretKey)}
}

// GenToken 生成 JWT
func (j *MyJWT) GenToken(userId int, role string, uuid string) (string, error) {
	// 创建一个 MyClaims 实例
	claims := MyClaims{
		UserId: userId,
		Role:   role,
		UUID:   uuid,
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.Conf.JWT.Issuer,
			ExpiresAt: time.Now().Add(time.Duration(config.Conf.JWT.ExpireTime)).Unix(),
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的 secret 签名并获得完整的编码后的字符串 token
	return token.SignedString(j.SecretKey)
}

// ParseToken 解析 JWT
func (j *MyJWT) ParseToken(tokenString string) (*MyClaims, error) {
	// 解析 token
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
				return nil, TokenExpired
			} else if ve.Errors&(jwt.ValidationErrorNotValidYet) != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	// 校验 token
	if token.Valid {
		return claims, nil
	}

	return nil, TokenInvalid
}
