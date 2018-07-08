package api

import (
	"encoding/json"
	"fathom/core"
	"net/http"
)

// URL: /api/countries
var GetCountriesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	period := getRequestedPeriod(r)

	// get total
	stmt, err := core.DB.Prepare(`
		SELECT
		COUNT(DISTINCT(ip_address))
		FROM visits
		WHERE timestamp >= DATE_SUB(CURRENT_TIMESTAMP, INTERVAL ? DAY)`)
	checkError(err)
	defer stmt.Close()
	var total float32
	stmt.QueryRow(period).Scan(&total)

	// get rows
	stmt, err = core.DB.Prepare(`
		SELECT
		country,
		COUNT(DISTINCT(ip_address)) AS count
		FROM visits
		WHERE timestamp >= DATE_SUB(CURRENT_TIMESTAMP, INTERVAL ? DAY) AND country IS NOT NULL
		GROUP BY country
		ORDER BY count DESC
		LIMIT ?`)
	checkError(err)
	defer stmt.Close()
	rows, err := stmt.Query(period, defaultLimit)
	checkError(err)
	defer rows.Close()

	results := make([]Datapoint, 0)
	for rows.Next() {
		var d Datapoint
		err = rows.Scan(&d.Label, &d.Count)
		checkError(err)

		d.Percentage = float32(d.Count) / total * 100
		results = append(results, d)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
})
