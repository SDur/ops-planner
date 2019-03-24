package model

type db interface {
	SelectMembers() ([]*Member, error)
	InsertMember(member *Member) error
}
