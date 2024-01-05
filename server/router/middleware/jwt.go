package middleware

import (
	"blog-server/utils"
	"blog-server/utils/r"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 定义存放位置（Header 中的 Authorization）和格式（`Bearer [tokenString]`）
		token := context.Request.Header.Get("Authorization")
		// 判断是否为空
		if token == "" {
			r.SendCode(context, r.ERROR_TOKEN_NOT_EXIST)
			context.Abort()
			return
		}
		// 分隔 token，判断格式是否正确
		parts := strings.Split(token, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			r.SendCode(context, r.ERROR_TOKEN_TYPE_WRONG)
			context.Abort()
			return
		}
		// 解析 token
		claims, err := utils.GetJWT().ParseToken(parts[1])
		if err != nil {
			r.SendCodeWithData(context, r.ERROR_TOKEN_WRONG, err.Error())
			context.Abort()
			return
		}
		// 判断 token 是否过期
		if time.Now().Unix() > claims.ExpiresAt {
			r.SendCode(context, r.ERROR_TOKEN_RUNTIME)
			context.Abort()
			return
		}
		// 将请求信息保存到上下文 context，后续处理函数可以通过 context.Get("xxx") 获取对应的用户信息。
		context.Set("user_id", claims.UserId)
		context.Set("role", claims.Role)
		context.Set("uuid", claims.UUID)
		context.Next()
	}
}
