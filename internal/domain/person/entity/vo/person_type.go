package valueobjcet

type PersonType uint

const (
	INTERNAL PersonType = 1 + iota
	EXTERNAL
)
