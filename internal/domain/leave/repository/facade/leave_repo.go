package facade

import "github.com/dapr-ddd-action/internal/domain/leave/repository/po"

type LeaveRepository interface {
	Save(leavePO *po.LeavePO) error
	SaveEvent(leaveEventPO *po.LeaveEventPO) error
	FindById(id string) (po.LeavePO, error)
	QueryByApplicantId(applicantId string) ([]*po.LeavePO, error)
	QueryByApproverId(approverId string) ([]*po.LeavePO, error)
}
