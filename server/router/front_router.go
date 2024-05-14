package router

import (
	"blog-server/config"
	"blog-server/router/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FrontRouter 前台路由
func FrontRouter() http.Handler {
	router := gin.New()

	router.Use(middleware.Logger())
	// 基于 cookies 存储 session
	store := cookie.NewStore([]byte(config.Conf.Session.Salt))
	store.Options(sessions.Options{MaxAge: config.Conf.Session.ExpireTime})

	router.Use(sessions.Sessions(config.Conf.Session.Name, store))

	base := router.Group("api/front")
	{
		base.GET("/home")
		base.POST("/login", userAuthAPI.Login) // 前台登录

	}
	return router
}
