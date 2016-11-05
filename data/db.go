package data

import (
	_ "database/sql" // Database.
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

const (
	dialect  = "sqlite3"
	database = "test.db"
)

// ConnectDB Connects to sqlite3 database
func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(dialect, database)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to DB.")
	return db, err
}
