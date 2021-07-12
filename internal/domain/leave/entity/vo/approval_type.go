package vo

type ApprovalType uint

const (
	AGREE ApprovalType = 1 + iota
	REJECT
)
