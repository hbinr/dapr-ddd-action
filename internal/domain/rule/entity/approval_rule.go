package entity

import (
	"strconv"

	"github.com/dapr-ddd-action/internal/domain/leave/entity"
)

// ApprovalRule 审核规则
type ApprovalRule struct {
	PersonType     string
	LeaveType      string
	Duration       int64
	MaxLeaderLevel int
}

func (r ApprovalRule) GetByLeave(leave *entity.Leave) {
	r.PersonType = leave.Applicant.PersonType
	r.LeaveType = strconv.Itoa(int(leave.LeaveType))
	r.Duration = leave.Duration

}
