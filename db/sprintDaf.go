package db

import (
	"database/sql"
	"github.com/SDur/ops-planner/model"
	"github.com/lib/pq"
	"time"
)

func (p *pgDb) SelectCurrentSprint() (*model.Sprint, error) {
	row := p.dbConn.QueryRowx("SELECT * FROM sprints order by start desc LIMIT 1")
	var days []sql.NullInt64
	var id int64
	var nr int64
	var start time.Time
	if err := row.Scan(&id, &nr, &start, pq.Array(&days)); err != nil {
		return nil, err
	}
	var convertedDays [10]int64

	for i, d := range days {
		convertedDays[i] = int64(d.Int64)
	}

	s := &model.Sprint{
		Id:    id,
		Nr:    nr,
		Start: start,
		Days:  convertedDays,
	}
	return s, nil
}

func (p *pgDb) UpdateSprint(sprint *model.Sprint) error {
	_, e := p.dbConn.Exec("UPDATE sprints SET days = $1 WHERE nr = $2",
		pq.Array(sprint.Days),
		sprint.Nr)
	return e
}

func (p *pgDb) InsertSprint(sprint *model.Sprint) error {
	_, e := p.dbConn.Exec("INSERT into sprints (nr, start, days) values ($1, $2, $3)",
		sprint.Nr,
		sprint.Start,
		pq.Array(sprint.Days))
	return e
}
