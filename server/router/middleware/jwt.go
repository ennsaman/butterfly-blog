package middleware

import (
	"blog-server/utils"
	"blog-server/utils/r"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 定义存放位置（Header 中的 Authorization）和格式（`Bearer [tokenString]`）
		token := ctx.Request.Header.Get("Authorization")
		// 判断是否为空
		if token == "" {
			r.SendCode(ctx, r.ERROR_TOKEN_NOT_EXIST)
			ctx.Abort()
			return
		}
		// 分隔 token，判断格式是否正确
		parts := strings.Split(token, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			r.SendCode(ctx, r.ERROR_TOKEN_TYPE_WRONG)
			ctx.Abort()
			return
		}
		// 解析 token
		claims, err := utils.GetJWT().ParseToken(parts[1])
		if err != nil {
			r.SendCodeWithData(ctx, r.ERROR_TOKEN_WRONG, err.Error())
			ctx.Abort()
			return
		}
		// 判断 token 是否过期
		if time.Now().Unix() > claims.ExpiresAt {
			r.SendCode(ctx, r.ERROR_TOKEN_RUNTIME)
			ctx.Abort()
			return
		}
		// 将请求信息保存到上下文 ctx，后续处理函数可以通过 ctx.Get("xxx") 获取对应的用户信息。
		ctx.Set("user_id", claims.UserId)
		ctx.Set("role", claims.Role)
		ctx.Set("uuid", claims.UUID)
		ctx.Next()
	}
}
