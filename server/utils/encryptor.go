package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

var Encryptor *_encryptor

type _encryptor struct{}

// BcryptCheck 检查密码是否正确
func (*_encryptor) BcryptCheck(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// MD5 加密
func (*_encryptor) MD5(str string) string {
	// 创建 MD5 实例
	hash := md5.New()
	// 将字符串转换为字节数组并写入 MD5 实例
	hash.Write([]byte(str))
	// 获取计算得到的 MD5 哈希值
	hashInBytes := hash.Sum(nil)
	// 将字节数组转换为十六进制字符串
	md5String := hex.EncodeToString(hashInBytes)
	return md5String
}
