package db

import (
	"fmt"
	"github.com/SDur/ops-planner/model"
	"log"
)

func (p *pgDb) SelectMembers() ([]*model.Member, error) {
	people := make([]*model.Member, 0)
	if err := p.sqlSelectMembers.Select(&people); err != nil {
		return nil, err
	}
	return people, nil
}

func (p *pgDb) SelectMember(id int64) (*model.Member, error) {
	row := p.dbConn.QueryRowx(fmt.Sprintf("SELECT * FROM members where id = %d", id))
	var member model.Member
	if err := row.StructScan(&member); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &member, nil
}

func (p *pgDb) InsertMember(newMember *model.Member) error {
	_, e := p.dbConn.Exec("INSERT INTO members (firstname, lastname) VALUES ($1, $2)", newMember.Firstname, newMember.Lastname)
	return e
}

func (p *pgDb) DeleteMember(id int) error {
	_, e := p.dbConn.Exec("DELETE FROM members where id=$1", id)
	return e
}
