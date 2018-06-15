package core

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func SetupDatabaseConnection() *sql.DB {
	// setup db connection
	var err error
	DB, err = sql.Open("mysql", "root:123456@/ana")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Open doesn't open a connection. Validate DSN data:
	err = DB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	return DB
}
