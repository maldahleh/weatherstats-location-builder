package scraper

import (
	"os"
	"strings"
	"time"

	"weatherstatsLocations/downloader"
	"weatherstatsLocations/reader"
	cs "weatherstatsLocations/scraper/station"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

const rootUrl = "http://dd.weatheroffice.ec.gc.ca/climate/observations/daily/csv/"
const timeoutMinutes = 20

var provinces = [...]string{
	"AB",
	"BC",
	"MB",
	"NB",
	"NL",
	"NS",
	"NT",
	"NU",
	"ON",
	"PE",
	"QC",
	"SK",
	"YT",
}

type climateStations map[string]*cs.ClimateStation
type ProvincialStations map[string]climateStations

func Scrape() ProvincialStations {
	provinceStations := make(ProvincialStations)
	for _, e := range provinces {
		provinceStations[e] = scrape(e)
	}

	return provinceStations
}

func scrape(province string) climateStations {
	url := rootUrl + province + "/"
	climateData := make(climateStations)

	c := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(0),
		colly.UserAgent("Mozilla/5.0"),
		colly.IgnoreRobotsTxt(),
	)

	c.SetRequestTimeout(timeoutMinutes * time.Minute)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		path := e.Attr("href")
		if !strings.Contains(path, ".csv") {
			return
		}

		split := strings.Split(path, "_")
		if len(split) != 6 {
			return
		}

		stationId := split[3]

		timestamp := split[4]
		timestampSplit := strings.Split(timestamp, "-")
		if len(timestampSplit) != 2 {
			return
		}

		month := timestampSplit[1]
		year := timestampSplit[0]

		station := climateData[stationId]
		if station == nil {
			station = cs.NewClimateStation()

			err := downloader.DownloadFile(path, url + path)
			if err != nil {
				station.Name = "N/A"
			} else {
				station.Name = reader.RetrieveStationName(path)

				fileError := os.Remove(path)
				if fileError != nil {
					log.Errorln("Failed to delete:", path, "\nError:", fileError)
				}
			}
		}

		availableData := station.AvailableData
		yearData := availableData[year]
		if yearData == nil {
			yearData = []string{}
		}

		yearData = append(yearData, month)
		availableData[year] = yearData

		climateData[stationId] = station
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Errorln("Request URL:", r.Request.URL, "\nError:", err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Errorln("Failed to visit:", url, "\nError:", err.Error())
	}

	c.Wait()
	return climateData
}
