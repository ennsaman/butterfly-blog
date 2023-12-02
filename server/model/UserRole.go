package model

// 用户角色关系模型
type UserRole struct {
	UserId int64 `json:"user_id"`
	RoleId int64 `json:"role_id"`
}
