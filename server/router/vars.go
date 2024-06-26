package router

import (
	"blog-server/api/front"
	v1 "blog-server/api/v1"
)

// 后台接口
var (
	userAPI     v1.User     // 用户api
	userAuthAPI v1.UserAuth // 用户认证api
)

// 前台接口
var (
	frontBlogInfoAPI front.BlogInfo // 前台博客信息
)
