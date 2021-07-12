package valueobjcet

type PersonStatus uint

const (
	ENABLE PersonStatus = 1 + iota
	DISABLE
)
