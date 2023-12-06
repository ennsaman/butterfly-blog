package v1

import "github.com/gin-gonic/gin"

type UserAuth struct{}

// 注册
func (*UserAuth) Register(context *gin.Context) {

}

// 登录
func (*UserAuth) Login(context *gin.Context) {

}

// 退出登录
func (*UserAuth) Logout(context *gin.Context) {

}
