package resp

import "time"

type LoginVo struct {
	Id         int    `json:"id"`
	UserInfoId int    `json:"user_info_id"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Intro      string `json:"intro"`
	Website    string `json:"website"`

	LoginType     int       `json:"login_type"`
	LastLoginTime time.Time `json:"last_login_time"`
	IpAddress     string    `json:"ip_address"`
	IpSource      string    `json:"ip_source"`

	Token string `json:"token"`
}

// UserVO 后台列表 VO
type UserVO struct {
	Id         int    `json:"id"`
	UserInfoId int    `json:"user_info_id"`
	Avatar     string `json:"avatar"`
	Nickname   string `json:"nickname"`

	LoginType     int       `json:"login_type"`
	LastLoginTime time.Time `json:"last_login_time"`
	IpAddress     string    `json:"ip_address"`
	IpSource      string    `json:"ip_source"`
	IsDisable     int       `json:"is_disable"`
	// Intro         string       `json:"intro"`
}
