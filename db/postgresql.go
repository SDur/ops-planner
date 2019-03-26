package db

import (
	"database/sql"
	"github.com/SDur/ops-planner/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func (p *pgDb) SelectCurrentSprint() *model.Sprint {
	row := p.dbConn.QueryRowx("SELECT * FROM sprints LIMIT 1")
	var s model.Sprint
	row.StructScan(&s)
	return &s
}
