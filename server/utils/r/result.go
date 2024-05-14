package r

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应结构体
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// ReturnJson 返回 JSON 数据
func ReturnJson(ctx *gin.Context, httpCode int, code int, msg string, data any) {
	// ctx.Header("", "") // 根据需要在头部添加其他信息
	ctx.JSON(httpCode, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Send 语法糖函数封装
func Send(ctx *gin.Context, httpCode int, code int, data any) {
	ReturnJson(ctx, httpCode, code, GetMsg(code), data)
}

// Success 成功响应，无数据
func Success(ctx *gin.Context) {
	Send(ctx, http.StatusOK, SUCCESS, nil)
}

// SuccessWithData 成功响应，有数据
func SuccessWithData(ctx *gin.Context, data any) {
	Send(ctx, http.StatusOK, SUCCESS, data)
}

// SendCode 根据 code 获取 msg，无数据
func SendCode(ctx *gin.Context, code int) {
	Send(ctx, http.StatusOK, code, nil)
}

// SendCodeWithData 根据 code 获取 msg，有数据
func SendCodeWithData(ctx *gin.Context, code int, data any) {
	Send(ctx, http.StatusOK, code, data)
}
