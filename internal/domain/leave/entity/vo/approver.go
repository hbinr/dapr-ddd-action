package vo

// Approver 审批人
type Approver struct {
	PersonId   string
	PersonName string
	Level      int
}

func (a *Approver) FromPerson() {

}
