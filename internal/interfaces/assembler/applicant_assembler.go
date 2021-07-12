package assembler

import (
	leaveVal "github.com/dapr-ddd-action/internal/domain/leave/entity/vo"
	"github.com/dapr-ddd-action/internal/interfaces/dto"
)

var Applicant = applicantAssembler{}

type applicantAssembler struct{}

func (a applicantAssembler) ToDTO(applicant *leaveVal.Applicant) *dto.ApplicantDTO {
	resdto := new(dto.ApplicantDTO)
	resdto.PersonId = applicant.PersonId
	resdto.PersonName = applicant.PersonName
	return resdto
}

func (a applicantAssembler) ToDO() {

}
