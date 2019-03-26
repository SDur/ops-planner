package db

func (p *pgDb) createTablesIfNotExist() error {
	create_members_sql := `

       CREATE TABLE IF NOT EXISTS members (
       id SERIAL NOT NULL PRIMARY KEY,
       firstname TEXT NOT NULL,
       lastname TEXT NOT NULL);

    `
	if rows, err := p.dbConn.Query(create_members_sql); err != nil {
		return err
	} else {
		rows.Close()
	}

	create_sprint_sql := `

       CREATE TABLE IF NOT EXISTS sprints (
       id SERIAL NOT NULL PRIMARY KEY,
       nr INTEGER NOT NULL,
       start DATE NOT NULL,
       days INTEGER ARRAY[10]);
    `

	if rows, err := p.dbConn.Query(create_sprint_sql); err != nil {
		return err
	} else {
		rows.Close()
	}

	return nil
}
