package query

type UsersPageQuery struct {
	CurrentPage int
	PageSize    int
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
