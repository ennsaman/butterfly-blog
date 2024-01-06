package req

// PageQuery 获取分页数据
type PageQuery struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	Keyword  string `form:"keyword"`
}
