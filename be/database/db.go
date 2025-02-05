package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabaseConn() (*Database, error) {
	driverName := "postgres"
	dataSource := "postgresql://root:24434@localhost:5432/go-websocket-sesat?sslmode=disable"
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &Database{
		db: db,
	}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
