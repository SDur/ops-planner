package model

import "time"

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

func (m *Model) AddMember(newMember *Member) error {
	return m.InsertMember(newMember)
}

func (m *Model) RemoveMember(id int) error {
	return m.DeleteMember(id)
}

func (m *Model) getMemberForToday() (*Member, error) {
	return m.GetMemberForDate(time.Now())
}

func (m *Model) CurrentSprint() (*Sprint, error) {
	return m.SelectCurrentSprint()
}

func (m *Model) SaveSprint(sprint *Sprint) error {
	return m.UpdateSprint(sprint)
}

func (m *Model) AddSprint(sprint *Sprint) error {
	return m.InsertSprint(sprint)
}
