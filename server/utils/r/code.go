package r

// 响应码
const (
	SUCCESS = 200 // 成功
	FAIL    = 400 // 失败
	ERROR   = 500 // 错误
)

// 响应信息
var MsgFlags = map[int]string{
	SUCCESS: "成功",
	FAIL:    "失败",
	ERROR:   "错误",
}

// 获取响应信息
func GetMsg(code int) string {
	return MsgFlags[code]
}
