package pagination

type BasePageQuery struct {
	CurrentPage int `query:"current_page"`
	PageSize    int `query:"page_size"`
}

func GetPageOffset(pageNum, pageSize int) int {
	return (pageNum - 1) * pageSize
}

func (up *BasePageQuery) GetPageSize() int {
	if up.PageSize == 0 {
		return 10
	}

	return up.PageSize
}

func (up *BasePageQuery) GetCurrentPage() int {
	if up.CurrentPage == 0 {
		return 1
	}
	return up.CurrentPage
}
