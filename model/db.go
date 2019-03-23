package model

type db interface {
	SelectMembers() ([]*Member, error)
}
