package vo

type LeaveType uint

const (
	INTERNAL LeaveType = 1 + iota
	EXTERNAL
	OFFICIAL
)
