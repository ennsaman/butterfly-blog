package service

import (
	"blog-server/model/resp"
	"github.com/gin-gonic/gin"
)

type User struct{}

// Login 登录
func (*User) Login(context *gin.Context, username, password string) (loginVo resp.LoginVo, code int) {

	return
}
