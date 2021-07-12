package dto

type ApprovalInfoDTO struct {
	ApprovalInfoId string
	Msg            string
	ApprovalTime   int64
	ApproverDTO
}
