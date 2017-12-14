package main

import (
	"database/sql"

	_ "dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
}
