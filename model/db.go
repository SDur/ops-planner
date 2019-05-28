package model

import "time"

type db interface {
	SelectMembers() ([]*Member, error)
	InsertMember(member *Member) error
	DeleteMember(id int) error
	GetMemberForDate(time time.Time) (*Member, error)
	SelectLatestSprint() (*Sprint, error)
	SelectSprints() ([]*Sprint, error)
	UpdateSprint(sprint *Sprint) error
	InsertSprint(sprint *Sprint) error
}
