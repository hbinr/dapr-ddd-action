package query

import "github.com/dapr-ddd-action/pkg/util/pagination"

type UsersPageQuery struct {
	pagination.BasePageQuery
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
