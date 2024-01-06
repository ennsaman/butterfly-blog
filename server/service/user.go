package service

import (
	"blog-server/dao"
	"blog-server/model"
	"blog-server/model/dto"
	"blog-server/model/req"
	"blog-server/model/resp"
	"blog-server/utils"
	"blog-server/utils/r"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type User struct{}

// Login 登录 TODO
func (*User) Login(ctx *gin.Context, username, password string) (loginVo resp.LoginVo, code int) {
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
	userDetailDTO := convertUserDetailDTO(userAuth, ctx)

	// 生成 UUID
	uuid := utils.Encryptor.MD5(userDetailDTO.IpAddress)
	// 生成 Token
	token, err := utils.GetJWT().GenToken(userAuth.Id, "test", uuid)
	if err != nil {
		return resp.LoginVo{}, r.ERROR_TOKEN_CREATE
	}
	userDetailDTO.Token = token

	// 更新用户登录信息
	dao.Update(&model.UserAuth{
		Id:        userAuth.Id,
		IpAddress: userDetailDTO.IpAddress,
		IpSource:  userDetailDTO.IpSource,
	}, "ip_address", "ip_source")

	// 保存用户信息到 Session 和 Redis
	session := sessions.Default(ctx)
	sessionInfoStr := utils.Json.Marshal(dto.SessionInfo{UserDetailDTO: userDetailDTO})
	session.Set(KEY_USER+uuid, sessionInfoStr)
	err = session.Save()
	if err != nil {
		utils.Logger.Error("保存用户信息到 Session 中失败: ", zap.Error(err))
	}
	// 返回
	return userDetailDTO.LoginVo, r.SUCCESS
}

// Logout 登出
func (u *User) Logout(ctx *gin.Context) {
	// 获取 UUID
	uuid := utils.GetFromContext[string](ctx, "uuid")
	// 删除 Session 中的用户信息
	session := sessions.Default(ctx)
	session.Delete(KEY_USER + uuid)
	err := session.Save()
	if err != nil {
		utils.Logger.Error("删除 Session 中的用户信息失败: ", zap.Error(err))
	}
	// 删除 Redis 中的用户信息 TODO
}

// GetList 获取用户列表
func (*User) GetList(req req.GetUsers) resp.PageResult[[]resp.UserVO] {
	// 获取记录数
	count := userDao.GetCount(req)
	// 获取用户列表
	list := userDao.GetList(req)
	// 返回
	return resp.PageResult[[]resp.UserVO]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    count,
		List:     list,
	}
}

func convertUserDetailDTO(userAuth model.UserAuth, ctx *gin.Context) dto.UserDetailDTO {
	// 获取用户详细信息
	userInfoId := userAuth.UserInfoId
	userInfo := dao.GetOne(model.UserInfo{}, "id = ?", userInfoId)

	// 填入角色 TODO

	// 从 ctx 中获取用户登录相关信息 TODO

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
