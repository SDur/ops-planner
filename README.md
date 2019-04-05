# ops-planner
OPSer van de dag planner

een hobby project in Go met een React frontend

# setting up
- run een postgres db via docker: `docker run -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres:alpine`
- login de db: `psql -h localhost -U postgres -d postgres`
- creeer de db `opsdb`: `create database opsdb;`
- run `go run main.go` 
- met de volgende flag: `-db-connect "host=localhost dbname=opsdb user=postgres password=postgres sslmode=disable"`

na de eerste keer draaien zijn de tables gecreeerd, 
vervolgens kan er een sprint toegevoegd worden alsvolgt (met *psql*):
- `\c opsdb`
- `insert into sprints(nr, start, days) values(1, '2019-03-28', '{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}');`
