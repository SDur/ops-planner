package model

type db interface {
	SelectMembers() ([]*Member, error)
	InsertMember(member *Member) error
	DeleteMember(id int) error
}
