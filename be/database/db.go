package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabaseConn() (*Database, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("failed to load environment")
	}

	dataSource := os.Getenv("DATABASE_CONNECTION")

	db, err := sql.Open("postgres", dataSource)
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
