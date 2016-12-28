package data

import (
	_ "database/sql" // Database.
	"fmt"
	"log"
	"os"
	"streamcat-api/configuration"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver.
)

const dialect = "postgres"

var connectionString = os.Getenv("DATABASE")

func init() {
	config := configuration.ConfigurationSetup()
	if connectionString == "" {
		connectionString = config.Database.Database
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
