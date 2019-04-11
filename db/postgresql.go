package db

import (
	"database/sql"
	"github.com/SDur/ops-planner/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Config struct {
	ConnectString string
}

func InitDb(cfg Config) (*pgDb, error) {
	if dbConn, err := sqlx.Connect("postgres", cfg.ConnectString); err != nil {
		return nil, err
	} else {
		p := &pgDb{dbConn: dbConn}
		if err := p.dbConn.Ping(); err != nil {
			return nil, err
		}
		if err := p.createTablesIfNotExist(); err != nil {
			return nil, err
		}
		if err := p.prepareSqlStatements(); err != nil {
			return nil, err
		}
		return p, nil
	}
}

type pgDb struct {
	dbConn *sqlx.DB

	sqlSelectMembers *sqlx.Stmt
	sqlInsertMember  *sqlx.NamedStmt
	sqlSelectMember  *sql.Stmt
}

func (p *pgDb) prepareSqlStatements() (err error) {

	if p.sqlSelectMembers, err = p.dbConn.Preparex(
		"SELECT id, firstname, lastname FROM members",
	); err != nil {
		return err
	}
	if p.sqlInsertMember, err = p.dbConn.PrepareNamed(
		"INSERT INTO members (firstname, lastname) VALUES (:firstname, :lastname) RETURNING id, firstname, lastname",
	); err != nil {
		return err
	}
	if p.sqlSelectMember, err = p.dbConn.Prepare(
		"SELECT id, firstname, lastname FROM members WHERE id = $1",
	); err != nil {
		return err
	}

	return nil
}

func (p *pgDb) GetMemberForDate(date time.Time) (*model.Member, error) {
	sprint, e := p.SelectCurrentSprint()
	if e != nil {
		return nil, e
	}
	days := date.Sub(sprint.Start).Hours() / 24
	log.Printf("Given day is %f away from start of spring %s", days, sprint.Start)

	startCopy := sprint.Start
	var sprintDay = 0

	for startCopy.Truncate(24 * time.Hour).Equal(date.Truncate(24 * time.Hour)) {
		if startCopy.Weekday() == time.Saturday || startCopy.Weekday() == time.Sunday {
			continue
		} else {
			sprintDay++
			startCopy.Add(1 * time.Hour)
		}
	}
	log.Printf("Day is %d nth day of the sprint", sprintDay)
	memberId := sprint.Days[sprintDay]
	log.Printf("Member id is %d", memberId)
	if memberId != -1 {
		member, e := p.SelectMember(memberId)
		if e != nil {
			log.Println(e.Error())
			return nil, e
		}
		log.Println("returning member")
		return member, nil
	}
	return nil, nil
}
