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
func ReturnJson(c *gin.Context, httpCode int, code int, msg string, data any) {
	// c.Header("", "") // 根据需要在头部添加其他信息
	c.JSON(httpCode, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Send 语法糖函数封装
func Send(c *gin.Context, httpCode int, code int, data any) {
	ReturnJson(c, httpCode, code, GetMsg(code), data)
}

// Success 成功响应，无数据
func Success(c *gin.Context) {
	Send(c, http.StatusOK, SUCCESS, nil)
}

// SuccessWithData 成功响应，有数据
func SuccessWithData(c *gin.Context, data any) {
	Send(c, http.StatusOK, SUCCESS, data)
}

// SendCode 根据 code 获取 msg，无数据
func SendCode(c *gin.Context, code int) {
	Send(c, http.StatusOK, code, nil)
}

// SendCodeWithData 根据 code 获取 msg，有数据
func SendCodeWithData(c *gin.Context, code int, data any) {
	Send(c, http.StatusOK, code, data)
}
