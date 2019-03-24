package model

type db interface {
	SelectMembers() ([]*Member, error)
	AddMember(newMember *Member) error
}
