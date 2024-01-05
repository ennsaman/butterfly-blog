package router

import (
	"blog-server/config"
	"blog-server/router/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BackRouter 后台路由
func BackRouter() http.Handler {
	router := gin.New()

	router.Use(middleware.Logger()) // 日志中间件

	// 使用 cookie 存储引擎
	store := cookie.NewStore([]byte(config.Conf.Session.Salt))
	// session 存储时间跟 JWT 过期时间一致
	store.Options(sessions.Options{MaxAge: int(config.Conf.JWT.ExpireTime) * 3600})
	router.Use(sessions.Sessions(config.Conf.Session.Name, store))

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
		auth.GET("/logout", userAuthAPI.Logout) // 退出登录
	}

	return router
}
