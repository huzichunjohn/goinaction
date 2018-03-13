package main

import (
	"shiori/cmd"
	db "shiori/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqliteDB, err := db.OpenSQLiteDatabase()
	checkError(err)

	cmd.DB = sqliteDB
	cmd.Execute()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
