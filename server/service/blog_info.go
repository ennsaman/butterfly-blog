package service

import (
	"blog-server/model/resp"
)

type BlogInfo struct {
}

// GetFrontHomeInfo 获取前台首页信息 TODO
func (*BlogInfo) GetFrontHomeInfo() resp.FrontBlogHomeVO {
	return resp.FrontBlogHomeVO{}
}
