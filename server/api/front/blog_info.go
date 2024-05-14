package front

import (
	"blog-server/utils/r"
	"github.com/gin-gonic/gin"
)

type BlogInfo struct{}

func (*BlogInfo) GetFrontHomeInfo(ctx *gin.Context) {
	r.SuccessWithData(ctx, blogInfoService.GetFrontHomeInfo)
}
