package db

import "github.com/SDur/ops-planner/model"

func (p *pgDb) SelectMembers() ([]*model.Member, error) {
	people := make([]*model.Member, 0)
	if err := p.sqlSelectMembers.Select(&people); err != nil {
		return nil, err
	}
	return people, nil
}

func (p *pgDb) InsertMember(newMember *model.Member) error {
	_, e := p.dbConn.Exec("INSERT INTO members (firstname, lastname) VALUES ($1, $2)", newMember.Firstname, newMember.Lastname)
	return e
}

func (p *pgDb) DeleteMember(id int) error {
	_, e := p.dbConn.Exec("DELETE FROM members where id=$1", id)
	return e
}
