package query

type UsersPageQuery struct {
	CurrentPage int `query:"current_page"`
	PageSize    int `query:"page_size"`
}

func (up UsersPageQuery) GetPageSize() int {
	if up.PageSize == 0 {
		return 10
	}

	return up.PageSize
}

func (up UsersPageQuery) GetCurrentPage() int {
	if up.CurrentPage == 0 {
		return 1
	}
	return up.CurrentPage
}
