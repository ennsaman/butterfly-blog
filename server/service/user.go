package service

import (
	"blog-server/dao"
	"blog-server/model"
	"blog-server/model/resp"
	"blog-server/utils/r"
	"github.com/gin-gonic/gin"
)

type User struct{}

// Login 登录 TODO
func (*User) Login(context *gin.Context, username, password string) (loginVo resp.LoginVo, code int) {
	// 获取用户认证信息
	userAuth := dao.GetOne(model.UserAuth{}, "username = ?", username)
	if userAuth.Id == 0 {
		return loginVo, r.ERROR_USER_NOT_EXIST
	}
	return
}
