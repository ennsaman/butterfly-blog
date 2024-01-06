package service

import "blog-server/dao"

// redis key
const (
	KEY_CODE = "code:" // 验证码
	KEY_USER = "user:" // 记录用户

)

var (
	userDao dao.User
)
