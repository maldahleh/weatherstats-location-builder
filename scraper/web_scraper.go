package scraper

import (
	"net/http"
	"os"
	"strings"

	"weatherstatsLocations/downloader"
	"weatherstatsLocations/reader"
	cs "weatherstatsLocations/scraper/station"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

const rootUrl = "https://dd.weather.gc.ca/climate/observations/daily/csv/"

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

	res, err := http.Get(url)
	if err != nil {
		log.Error("couldn't retrieve url", url, "error", err)
		return climateData
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Error("status code error", res.StatusCode, res.Status)
		return climateData
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Error("failed to load document", err)
	}

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		path, _ := s.Attr("href")
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

	return climateData
}
