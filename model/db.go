package model

type db interface {
	selectMembers() ([]*Member, error)
}
