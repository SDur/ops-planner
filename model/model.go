package model

type Model struct {
	db
}

func New(db db) *Model {
	return &Model{
		db: db,
	}
}

func (m *Model) Members() ([]*Member, error) {
	return m.SelectMembers()
}

func (m *Model) AddMember(newMember Member) error {
	return m.AddMember(newMember)
}
