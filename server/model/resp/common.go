package resp

type PageResult[T any] struct {
	PageSize int   `json:"pageSize"`
	PageNum  int   `json:"pageNum"`
	Total    int64 `json:"total"`
	List     T     `json:"pageData"` // ! 注意这里的别名
}
