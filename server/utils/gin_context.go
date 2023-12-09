package utils

import (
	"blog-server/utils/r"
	"github.com/gin-gonic/gin"
	"log"
)

// BindJSON JSON 绑定
func BindJSON[T any](context *gin.Context) (data T) {
	err := context.ShouldBindJSON(&data)
	if err != nil {
		log.Fatal("BindJSON：", err)
	}
	return
}

// Validate 参数合法性校验
func Validate(context *gin.Context, data any) {
	validateAns := Validator.Validate(data)
	if validateAns != "" {
		r.ReturnJson(context, 200, 400, validateAns, nil)
		panic(nil)
	}
}

// BindJSONAndValid JSON 绑定验证 + 合法性校验
func BindJSONAndValid[T any](context *gin.Context) (data T) {
	err := context.ShouldBindJSON(&data)
	if err != nil {
		log.Fatal("BindJSON：", err)
	}
	Validate(context, &data)
	return data
}
