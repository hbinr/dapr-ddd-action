package po

import (
	"time"

	"github.com/dapr-ddd-action/internal/domain/leave/event"
)

// LeaveEventPO 审批信息
type LeaveEventPO struct {
	Id             int64
	LeaveEventType event.LeaveEventType
	Source         string // uuid
	Data           string
	Date           time.Time
}
