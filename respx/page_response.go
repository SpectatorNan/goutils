package respx

import "math"

type PageResp struct {
	PageData
	Data interface{} `json:"list"`
}

type PageData struct {
	Page       uint `json:"page"`
	Size       uint `json:"size"`
	TotalPage  uint `json:"totalPage"`
	TotalCount uint `json:"totalCount"`
	IsLastPage bool `json:"isLastPage"`
}

func (p PageResp) Response() *Response {
	return NewSuccessResponse(p)
}

func NewPageResp(page, pageSize int, total int64, data interface{}) *PageResp {
	pageData := createPaging(page, pageSize, total)
	return &PageResp{
		PageData: pageData,
		Data:     data,
	}
}
func createPaging(page int, pagesize int, total int64) PageData {
	if page < 1 {
		page = 1
	}
	if pagesize < 1 {
		pagesize = 10
	}

	pageCount := math.Ceil(float64(total) / float64(pagesize))

	lastPage := false
	if page >= int(pageCount) {
		//page = int(pageCount)
		lastPage = true
	}

	paging := PageData{}
	paging.Page = uint(page)
	paging.Size = uint(pagesize)
	paging.TotalPage = uint(pageCount)
	paging.IsLastPage = lastPage
	paging.TotalCount = uint(total)
	return paging
}
