package scraper

import (
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"weatherstatsLocations/downloader"
	"weatherstatsLocations/reader"
	cs "weatherstatsLocations/scraper/station"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

const rootUrl string = "http://dd.weatheroffice.ec.gc.ca/climate/observations/daily/csv/"

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

func ScrapeProvinces() ProvincialStations {
	provinceStations := make(ProvincialStations)
	for _, e := range provinces {
		provinceStations[e] = scrape(e)
	}

	return provinceStations
}

func scrape(province string) climateStations {
	climateData := make(climateStations)

	c := colly.NewCollector(
		colly.AllowedDomains(
			"dd.weatheroffice.ec.gc.ca",
			"www.dd.weatheroffice.ec.gc.ca",
		),
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.MaxDepth(0),
		colly.UserAgent("Mozilla/5.0"),
	)

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ResponseHeaderTimeout: 120 * time.Second,
	})

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

			err := downloader.DownloadFile(path, rootUrl + province + "/" + path)
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
		log.Errorln("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(rootUrl + province + "/")
	if err != nil {
		log.Errorln("Failed to visit:", rootUrl, "with error:", err.Error())
	}

	c.Wait()
	return climateData
}
