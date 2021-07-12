package event

import (
	"encoding/json"
	"time"

	"github.com/dapr-ddd-action/internal/domain/leave/entity"
	"github.com/dapr-ddd-action/internal/infrastructure/common/event"
)

type LeaveEvent struct {
	event.DomainEvent
	LeaveEventType
}

func (l *LeaveEvent) Create(eventType LeaveEventType, leave entity.Leave) error {
	l.Id = "1" // uuid
	l.LeaveEventType = eventType
	l.CreateTime = time.Now()

	leaveStr, err := json.Marshal(leave)
	if err != nil {
		return err
	}

	l.Data = string(leaveStr)
	return nil
}
