package event

type LeaveEventType uint

const (
	CREATE_EVENT LeaveEventType = 1 + iota
	AGREE_EVENT
	REJECT_EVENT
	APPROVED_EVENT
)
