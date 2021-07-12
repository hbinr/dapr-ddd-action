package entity

import (
	"time"

	leaveVO "github.com/dapr-ddd-action/internal/domain/leave/entity/vo"
)

type Leave struct {
	Id                   string
	Applicant            *leaveVO.Applicant
	Approver             *leaveVO.Approver
	LeaveType            leaveVO.LeaveType
	Status               leaveVO.LeaveStatus
	StartTime            time.Time
	EndTime              time.Time
	Duration             int64
	LeaderMaxLevel       int // 审批领导的最大级别
	currentApprovalInfo  *ApprovalInfo
	historyApprovalInfos []*ApprovalInfo
}

func (l *Leave) GetDuration() int64 {
	return l.EndTime.Unix() - l.StartTime.Unix()
}

func (l *Leave) Create() {
	l.Status = leaveVO.APPROVING
	l.StartTime = time.Now()
}

func (l *Leave) Agree(nextApprover *leaveVO.Approver) {
	l.Status = leaveVO.APPROVING
	l.Approver = nextApprover
}

func (l *Leave) Reject(approver *leaveVO.Approver) {
	l.Status = leaveVO.REJECTED
	l.Approver = approver
}

func (l *Leave) Finish() {
	l.Status = leaveVO.APPROVED
	l.Approver = nil
	l.EndTime = time.Now()
	l.Duration = l.EndTime.Unix() - l.StartTime.Unix()
}

func (l *Leave) AddHistoryApprovalInfo(approvalInfo *ApprovalInfo) {
	if l.historyApprovalInfos == nil {
		l.historyApprovalInfos = make([]*ApprovalInfo, 1)
	}
	l.historyApprovalInfos = append(l.historyApprovalInfos, approvalInfo)
}
