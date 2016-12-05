package data

import (
	_ "database/sql" // Database.
	"fmt"
	"log"
	"os"
	"stream-api/configuration"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

const dialect = "sqlite3"

var database = os.Getenv("DATABASE")

func init() {
	config := configuration.ConfigurationSetup()
	if database == "" {
		database = config.Database.Database
	}
}

// ConnectDB Connects to sqlite3 database
func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(dialect, database)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to DB.")
	return db, err
}
