package main

import (
	"fathom/api"
	"fathom/core"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// TODO: Use Gorilla Mux router.
// TODO: Authentication.

func main() {
	// test .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := core.SetupDatabaseConnection()
	defer db.Close()

	r := mux.NewRouter()

	// register routes
	r.HandleFunc("/collect", api.CollectHandler).Methods("GET")
	r.Handle("/api/session", api.Login).Methods("POST")
	r.Handle("/api/session", api.Logout).Methods("DELETE")
	r.Handle("/api/visits/count/day", api.Authorize(api.GetVisitsDayCountHandler)).Methods("GET")
	r.Handle("/api/visits/count/realtime", api.Authorize(api.GetVisitsRealtimeCount)).Methods("GET")
	r.Handle("/api/visits", api.Authorize(api.GetVisitsHandler)).Methods("GET")
	r.Handle("/api/pageviews/count/day", api.Authorize(api.GetPageviewsDayCountHandler)).Methods("GET")
	r.Handle("/api/pageviews", api.Authorize(api.GetPageviewsHandler)).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))
}
