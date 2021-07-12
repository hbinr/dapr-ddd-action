package vo

type LeaveStatus uint

const (
	APPROVING LeaveStatus = 1 + iota
	APPROVED
	REJECTED
)
