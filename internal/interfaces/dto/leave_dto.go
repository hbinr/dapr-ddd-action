package dto

type LeaveDTO struct {
	LeaveId   string
	LeaveType string
	Status    string
	StartTime string
	EndTiem   string
	Duration  int64
	ApplicantDTO
	ApproverDTO
	CurrentApprovalInfoDTO     *ApprovalInfoDTO
	historyApprovalInfoDTOList []*ApprovalInfoDTO
}
