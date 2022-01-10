package pagination

func GetPageOffset(pageNum, pageSize int) int {
	return (pageNum - 1) * pageSize
}
