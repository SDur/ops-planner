package model

type db interface {
	SelectMembers() ([]*Member, error)
	InsertMember(member *Member) error
	DeleteMember(id int) error
	SelectCurrentSprint() (*Sprint, error)
	UpdateSprint(sprint *Sprint) error
	InsertSprint(sprint *Sprint) error
}
