package main

import (
	"database/sql"
	"fmt"
	"log"
	app "widgetapp"
	"widgetapp/http"
	"widgetapp/psql"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "huzichun"
	password = "123456"
	dbname   = "widget_demo"
)

var (
	userService   app.UserService
	widgetService app.WidgetService
)

func main() {
	// setup the DB connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	userService = &psql.UserService{DB: db}
	widgetService = &psql.WidgetService{DB: db}
	server := http.NewServer(userService, widgetService)
	log.Fatal(http.ListenAndServe(":3000", server))
}
