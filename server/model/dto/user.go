package dto

import "blog-server/model/resp"

type SessionInfo struct {
	UserDetailDTO
}

type UserDetailDTO struct {
	resp.LoginVo
	Password string `json:"password"`
	Role     string `json:"role"`
}
