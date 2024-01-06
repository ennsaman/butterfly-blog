package req

// GetUsers 获取用户列表
type GetUsers struct {
	PageQuery
	LoginType int8   `form:"login_type"`
	Nickname  string `form:"nickname"`
}

// Login 登录
type Login struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}
