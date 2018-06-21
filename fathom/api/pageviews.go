package api

import (
	"encoding/json"
	"fathom/core"
	"fathom/models"
	"net/http"
)

// URL: /api/pageviews
var GetPageviewsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	stmt, err := core.DB.Prepare(`SELECT
			path,
			COUNT(ip_address) AS pageviews,
			COUNT(DISTINCT(ip_address)) AS pageviews_unique
		FROM visits
		GROUP BY path`)
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	checkError(err)

	results := make([]models.Pageview, 0)
	defer rows.Close()
	for rows.Next() {
		var p models.Pageview
		err = rows.Scan(&p.Path, &p.Count, &p.CountUnique)
		checkError(err)
		results = append(results, p)
	}

	err = rows.Err()
	checkError(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
})
