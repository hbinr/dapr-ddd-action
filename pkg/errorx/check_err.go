package errorx

func IsNotFound(err error) bool {
	return From(err).Code == 404
}
