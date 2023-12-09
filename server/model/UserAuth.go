package model

import "time"

// UserAuth 用户身份信息模型
type UserAuth struct {
	Id            int64     `json:"id"`
	UserInfoId    string    `json:"user_info_id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	LoginType     string    `json:"login_type"`
	LastLoginTime string    `json:"last_login_time"`
	IpAddress     string    `json:"ip_address"`
	IpSource      string    `json:"ip_source"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}
