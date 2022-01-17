package errorx

func IsRecordNotFound(err error) bool {
	e := From(err)
	return e.Code == 404
}
