package service

import (
	"blog-server/dao"
	"blog-server/model"
	"blog-server/model/dto"
	"blog-server/model/resp"
	"blog-server/utils"
	"blog-server/utils/r"
	"github.com/gin-gonic/gin"
)

type User struct{}

// Login 登录 TODO
func (*User) Login(context *gin.Context, username, password string) (loginVo resp.LoginVo, code int) {
	// 获取用户认证信息
	userAuth := dao.GetOne(model.UserAuth{}, "username = ?", username)
	if userAuth.Id == 0 {
		return loginVo, r.ERROR_USER_NOT_EXIST
	}
	// 检查密码是否正确
	if !utils.Encryptor.BcryptCheck(password, userAuth.Password) {
		// 密码错误
		return loginVo, r.ERROR_PASSWORD_WRONG
	}

	// 获取用户详细信息
	userDetailDTO := convertUserDetailDTO(userAuth, context)

	// 生成 Token
	token, err := utils.GetJWT().GenToken(userAuth.Id, "test")
	if err != nil {
		return resp.LoginVo{}, r.ERROR_TOKEN_CREATE
	}
	userDetailDTO.Token = token

	// 更新用户登录信息

	// 保存用户信息到 Session 和 Redis

	// 返回
	return userDetailDTO.LoginVo, r.SUCCESS
}

func convertUserDetailDTO(userAuth model.UserAuth, context *gin.Context) dto.UserDetailDTO {
	// 获取用户详细信息
	userInfoId := userAuth.UserInfoId
	userInfo := dao.GetOne(model.UserInfo{}, "id = ?", userInfoId)

	// 填入角色 TODO

	// 从 context 中获取用户登录相关信息 TODO

	return dto.UserDetailDTO{
		LoginVo: resp.LoginVo{
			Id:            userAuth.Id,
			UserInfoId:    userInfoId,
			Username:      userAuth.Username,
			Nickname:      userInfo.Nickname,
			Avatar:        userInfo.Avatar,
			Email:         userInfo.Email,
			Intro:         userInfo.Intro,
			Website:       userInfo.Website,
			LoginType:     userAuth.LoginType,
			LastLoginTime: userAuth.LastLoginTime,
			IpAddress:     userAuth.IpAddress,
			IpSource:      userAuth.IpSource,
		},
		Password: userAuth.Password,
	}
}
