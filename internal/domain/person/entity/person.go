package entity

import (
	"time"

	personVO "github.com/dapr-ddd-action/internal/domain/person/entity/vo"
)

type Person struct {
	PersonId      string
	PersonName    string
	PersonType    personVO.PersonType
	PersonStatus  personVO.PersonStatus
	RoleLevel     int
	CreateTime    time.Time
	UpdateTime    time.Time
	Relationships []*Relationship
}

func (p *Person) Create() {
	p.CreateTime = time.Now()
	p.PersonStatus = personVO.ENABLE
}

func (p *Person) Enable() {
	p.UpdateTime = time.Now()
	p.PersonStatus = personVO.ENABLE
}

func (p *Person) Disable() {
	p.UpdateTime = time.Now()
	p.PersonStatus = personVO.DISABLE
}
