package event

import "time"

type DomainEvent struct {
	Id         string
	Source     string
	Data       string
	CreateTime time.Time
	UpdateTime time.Time
}
