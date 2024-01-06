package v1

import (
	"blog-server/model/req"
	"blog-server/utils"
	"blog-server/utils/r"
	"github.com/gin-gonic/gin"
)

type UserAuth struct{}

// Register 注册
func (*UserAuth) Register(ctx *gin.Context) {

}

// Login 登录
func (*UserAuth) Login(ctx *gin.Context) {
	// 参数绑定校验 + 参数合法性校验
	loginReq := utils.BindJSONAndValid[req.Login](ctx)
	// 业务逻辑
	loginVo, code := userService.Login(ctx, loginReq.Username, loginReq.Password)
	// 返回数据
	r.SendCodeWithData(ctx, code, loginVo)
}

// Logout 退出登录
func (*UserAuth) Logout(ctx *gin.Context) {
	// 业务逻辑
	userService.Logout(ctx)
	// 返回数据
	r.Success(ctx)
}
