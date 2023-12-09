package v1

import (
	"blog-server/model/req"
	"blog-server/utils"
	"github.com/gin-gonic/gin"
)

type UserAuth struct{}

// 注册
func (*UserAuth) Register(context *gin.Context) {

}

// 登录
func (*UserAuth) Login(context *gin.Context) {
	// 参数绑定校验 + 参数合法性校验
	loginReq := utils.BindJSONAndValid[req.Login](context)
	// 业务逻辑
	userService.Login(context, loginReq.Username, loginReq.Password)
	// 返回数据

}

// 退出登录
func (*UserAuth) Logout(context *gin.Context) {

}
