package model

import "time"

// 用户信息模型
type UserInfo struct {
	Id         int64     `json:"id"`
	Nickname   string    `json:"nickname"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	Intro      string    `json:"intro"`
	Website    string    `json:"website"`
	IsDisable  bool      `json:"is_disable"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
