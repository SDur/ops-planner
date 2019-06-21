package db

import (
	"database/sql"
	"github.com/SDur/ops-planner/model"
	"github.com/lib/pq"
	"log"
	"time"
)

func (p *pgDb) SelectLatestSprint() (*model.Sprint, error) {
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

func (p *pgDb) SelectSprintForDate(date time.Time) (*model.Sprint, error) {
	row := p.dbConn.QueryRow("SELECT * FROM sprints WHERE start <= $1 ORDER BY start DESC LIMIT 1", date)
	var days []sql.NullInt64
	var id int64
	var nr int64
	var start time.Time
	if err := row.Scan(&id, &nr, &start, pq.Array(&days)); err != nil {
		log.Println(err)
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

func (p *pgDb) SelectSprints() ([]*model.Sprint, error) {
	rows, err := p.dbConn.Queryx("SELECT * FROM sprints")
	if err != nil {
		return nil, err
	}

	sprints := make([]*model.Sprint, 0)

	for rows.Next() {
		var days []sql.NullInt64
		var id int64
		var nr int64
		var start time.Time
		if err := rows.Scan(&id, &nr, &start, pq.Array(&days)); err != nil {
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

		sprints = append(sprints, s)
	}
	return sprints, nil
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

func (p *pgDb) DeleteSprint(id int) error {
	_, e := p.dbConn.Exec("DELETE FROM sprints where id=$1", id)
	return e
}
