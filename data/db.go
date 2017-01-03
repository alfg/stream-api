package data

import (
	_ "database/sql" // Database.
	"fmt"
	"log"
	"os"
	"streamcat-api/settings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver.
)

const dialect = "postgres"

var connectionString = os.Getenv("DATABASE")

func init() {
	if connectionString == "" {
		connectionString = settings.Get().Database.Database
	}
}

// ConnectDB Connects to sqlite3 database
func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(dialect, connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to DB.")
	return db, err
}
