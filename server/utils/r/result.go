package r

import "github.com/gin-gonic/gin"

// 响应结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 返回 JSON 数据
func ReturnJson(c *gin.Context, httpCode, code int, msg string, data interface{}) {
	// c.Header("", "") // 根据需要在头部添加其他信息
	c.JSON(httpCode, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
