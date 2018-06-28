package db

import (
	"fathom/core"
	"fathom/models"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
)

var browserNames = []string{
	"Chrome",
	"Firefox",
	"Safari",
	"Internet Explorer",
}

var paths = []string{
	"/",
	"/",
	"/contact",
	"/about",
	"/checkout",
}

var browserLanguages = []string{
	"en-US",
	"en-US",
	"nl-NL",
	"fr_FR",
	"de_DE",
	"es_ES",
}

var screenResolutions = []string{
	"2560x1440",
	"1920x1080",
	"360x640",
}

func Seed(n int) {
	// prepare statement for inserting data
	stmt, err := core.DB.Prepare(`INSERT INTO visits(
		browser_language,
		browser_name,
		browser_version,
		country,
		device_os,
		ip_address,
		path,
		referrer_url,
		screen_resolution,
		timestamp
		) VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? `)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// insert X random hits
	for i := 0; i < n; i++ {
		// generate random timestamp
		date, err := time.Parse("Monday 2 Jan 2006", randomdata.FullDate())
		if err != nil {
			log.Fatal(err)
		}
		timestamp := fmt.Sprintf("%s %d:%d:%d", date.Format("2006-01-02"), randInt(0, 24), randInt(0, 60), randInt(0, 60))

		visit := models.Visit{
			Path:             randSliceElement(paths),
			IpAddress:        randomdata.IpV4Address(),
			DeviceOS:         "Linux x86_64",
			BrowserName:      randSliceElement(browserNames),
			BrowserVersion:   "54.0.2840.100",
			BrowserLanguage:  randSliceElement(browserLanguages),
			ScreenResolution: randSliceElement(screenResolutions),
			Country:          randomdata.Country(randomdata.ThreeCharCountry),
			ReferrerUrl:      "",
			Timestamp:        timestamp,
		}

		_, err = stmt.Exec(
			visit.BrowserLanguage,
			visit.BrowserName,
			visit.BrowserVersion,
			visit.Country,
			visit.DeviceOS,
			visit.IpAddress,
			visit.Path,
			visit.ReferrerUrl,
			visit.ScreenResolution,
			visit.Timestamp,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func randSliceElement(slice []string) string {
	return slice[randInt(0, len(slice))]
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
