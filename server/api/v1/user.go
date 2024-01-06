package v1

import (
	"blog-server/model/req"
	"blog-server/utils"
	"blog-server/utils/r"
	"github.com/gin-gonic/gin"
)

type User struct{}

// GetUserInfo 获取用户信息
func (*User) GetUserInfo(ctx *gin.Context) {

}

// GetList 获取用户列表
func (*User) GetList(ctx *gin.Context) {
	// 参数绑定
	getListReq := utils.BindQuery[req.GetUsers](ctx)
	// 业务逻辑
	userVo := userService.GetList(getListReq)
	// 返回数据
	r.SuccessWithData(ctx, userVo)
}
