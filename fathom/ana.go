package main

import (
	"fathom/api"
	"fathom/core"
	"net/http"
)

// TODO: Use Gorilla Mux router.
// TODO: Authentication.

func main() {
	db := core.SetupDatabaseConnection()
	defer db.Close()

	// register routes
	api.RegisterRoutes()
	http.HandleFunc("/collect", api.CollectHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/"+r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}
