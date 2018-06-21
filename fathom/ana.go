package main

import (
	"fathom/api"
	"fathom/core"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// TODO: Use Gorilla Mux router.
// TODO: Authentication.

func main() {
	db := core.SetupDatabaseConnection()
	defer db.Close()

	r := mux.NewRouter()

	// register routes
	r.HandleFunc("/collect", api.CollectHandler).Methods("GET")
	r.HandleFunc("/api/visits/count/day", api.GetVisitsDayCountHandler).Methods("GET")
	r.HandleFunc("/api/visits/count/realtime", api.GetVisitsRealtimeCount).Methods("GET")
	r.HandleFunc("/api/visits", api.GetVisitsHandler).Methods("GET")
	r.HandleFunc("/api/pageviews", api.GetPageviewsHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))
}
