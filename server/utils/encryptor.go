package utils

import "golang.org/x/crypto/bcrypt"

var Encryptor *encryptor

type encryptor struct{}

// BcryptCheck 检查密码是否正确
func (*encryptor) BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
