package router

import (
	"blog-server/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 后台路由
func BackRouter() http.Handler {
	router := gin.New()

	// 无需鉴权接口
	base := router.Group("/api/v1")
	{
		base.POST("/login", userAuthAPI.Login) // 后台登录
	}

	// 需要鉴权接口
	auth := base.Group("")

	// 加入 JWT 鉴权中间件
	auth.Use(middleware.JWTAuth())

	{

	}

	return nil
}
