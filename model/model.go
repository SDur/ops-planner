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

func (m *Model) GetMemberForToday() (*Member, error) {
	return m.GetMemberForDate(time.Now())
}

func (m *Model) Sprints() ([]*Sprint, error) {
	return m.SelectSprints()
}

func (m *Model) LatestSprint() (*Sprint, error) {
	return m.SelectLatestSprint()
}

func (m *Model) SaveSprint(sprint *Sprint) error {
	return m.UpdateSprint(sprint)
}

func (m *Model) AddSprint(sprint *Sprint) error {
	return m.InsertSprint(sprint)
}

func (m *Model) RemoveSprint(id int) error {
	return m.DeleteSprint(id)
}
