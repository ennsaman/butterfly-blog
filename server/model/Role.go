package model

// Role 角色模型
type Role struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Label     string `json:"label"`
	IsDisable bool   `json:"is_disable"`
}
