package entity

import leaveVal "github.com/dapr-ddd-action/internal/domain/leave/entity/vo"

type ApprovalInfo struct {
	ApprovalInfoId string
	Msg            string
	ApprovalTime   int64
	Approver       leaveVal.Approver
	ApprovalType   leaveVal.ApprovalType
}
