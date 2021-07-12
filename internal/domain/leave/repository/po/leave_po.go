package po

import (
	"time"

	leaveVal "github.com/dapr-ddd-action/internal/domain/leave/entity/vo"
	personVal "github.com/dapr-ddd-action/internal/domain/person/entity/vo"
)

// LeavePO 请假PO
type LeavePO struct {
	Id                        string
	ApplicantId               string
	ApplicantName             string
	ApplicantType             personVal.PersonType
	ApproverId                string
	ApproverName              string
	LeaveType                 leaveVal.LeaveType
	LeaveStatus               leaveVal.LeaveStatus
	Duration                  int64
	StartTime                 time.Time
	EndTime                   time.Time
	HistoryApprovalInfoPOList []*ApprovalInfoPO
}
