package utils

import (
	"blog-server/utils/r"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// BindJSON JSON 绑定
func BindJSON[T any](context *gin.Context) (data T) {
	err := context.ShouldBindJSON(&data)
	if err != nil {
		Logger.Error("BindJSON：", zap.Error(err))
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
		Logger.Error("BindJSONAndValid：", zap.Error(err))
		panic(r.ERROR_REQUEST_PARAM)
	}
	Validate(context, &data)
	return data
}

// GetFromContext 从 context 中获取数据
func GetFromContext[T any](context *gin.Context, key string) (data T) {
	value, exists := context.Get(key)
	if !exists {
		panic(r.ERROR_TOKEN_NOT_EXIST)
	}
	return value.(T)
}

// BindQuery Param 绑定
func BindQuery[T any](ctx *gin.Context) (data T) {
	if err := ctx.ShouldBindQuery(&data); err != nil {
		Logger.Error("BindQuery：", zap.Error(err))
		panic(r.ERROR_REQUEST_PARAM)
	}
	// TODO 检查是否有 PageSize 或 PageQuery 字段，并处理其值
	return
}
