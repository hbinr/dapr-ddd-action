package po

// ApprovalInfoPO 审批信息
type ApprovalInfoPO struct {
	ApprovalInfoId string // uuid
	LeaveId        string
	ApplicantId    string
	ApproverId     string
	ApproverName   string
	Msg            string
	ApproverLevel  int
	ApprovalTime   int64
}
