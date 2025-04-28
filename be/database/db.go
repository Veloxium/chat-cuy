package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

type Database struct {
	db *sql.DB
}

func NewDatabaseConn() *Database {
	// load environment variable
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load environment")
	}
	dataSource := os.Getenv("DATABASE_CONNECTION")
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic("failed to connect the database")
	}
	return &Database{
		db: db,
	}
}

func (d *Database) Ping() error {
	err := d.db.Ping()
	return err
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
