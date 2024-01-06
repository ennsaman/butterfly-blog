package dao

import (
	"blog-server/model/req"
	"blog-server/model/resp"
)

type User struct{}

// GetCount 获取用户数量
func (*User) GetCount(req req.GetUsers) (count int64) {
	/*
		SELECT COUNT(1)
		FROM t_user_info ui LEFT JOIN t_user_auth ua
		ON ua.user_info_id = ui.id
		WHERE login_type = 1 AND nickname LIKE '%xxx%'
	*/

	tx := DB.Select("COUNT(1)").Table("t_user_info ui").
		Joins("LEFT JOIN t_user_auth ua ON ua.user_info_id = ui.id")
	if req.LoginType != 0 {
		tx.Where("login_type = ?", req.LoginType)
	}
	if req.Nickname != "" {
		tx.Where("nickname LIKE ?", "%"+req.Nickname+"%")
	}
	tx.Count(&count)
	return
}

// GetList 获取用户列表
func (*User) GetList(req req.GetUsers) (list []resp.UserVO) {
	/*
		SELECT ua.id, ua.user_info_id, ui.avatar, ui.nickname, ua.login_type, ua.last_login_time, ua.ip_address, ua.ip_source, ui.is_disable
		FROM t_user_auth ua
		LEFT JOIN t_user_info ui
		ON ua.user_info_id = ui.id
		WHERE ua.login_type = ?
		AND ui.nickname = ?
		LIMIT pageSize * (pageNum - 1), pageSize
	*/

	tx := DB.Select("ua.id, ua.user_info_id, ui.avatar, ui.nickname, ua.login_type, ua.last_login_time, ua.ip_address, ua.ip_source, ui.is_disable").
		Table("t_user_auth ua").
		Joins("LEFT JOIN t_user_info ui ON ua.user_info_id = ui.id")
	if req.LoginType != 0 {
		tx.Where("ua.login_type = ?", req.LoginType)
	}
	if req.Nickname != "" {
		tx.Where("ui.nickname LIKE ?", "%"+req.Nickname+"%")
	}
	tx.Offset(req.PageSize * (req.PageNum - 1)).Limit(req.PageSize)
	tx.Find(&list)
	return
}
