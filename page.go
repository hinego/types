package types

type PageRes struct {
	Total int64 `json:"total" dc:"总计记录"`
}
type PageReq struct {
	limit  int
	Page   int               `json:"page" dc:"页码"`
	Size   int               `json:"size" dc:"分页"`
	Sorter map[string]string `json:"sorter" dc:"排序规则"`
}

func (r *PageReq) Limit(num int) {
	r.limit = num
}
func (r *PageReq) Offset() int {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.Size < 10 {
		r.Size = 10
	}
	if r.limit == 0 {
		r.limit = 1000
	}
	if r.limit > 0 && r.limit < r.Size {
		r.Size = r.limit
	}
	return r.Size * (r.Page - 1)
}
