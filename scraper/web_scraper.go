package scraper

import (
	"fmt"
	"strings"


	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

const rootUrl string = "http://dd.weatheroffice.ec.gc.ca/climate/observations/daily/csv/"
var provinces = [...]string {
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

func Scrape(province string) {
	c := colly.NewCollector(
		colly.AllowedDomains(
			"dd.weatheroffice.ec.gc.ca",
			"www.dd.weatheroffice.ec.gc.ca",
		),
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.MaxDepth(1),
		colly.UserAgent("Mozilla/5.0"),
	)

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

		fmt.Println(stationId)
		fmt.Println(month)
		fmt.Println(year)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Errorln("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(rootUrl + province + "/")
	if err != nil {
		log.Errorln("Failed to visit:", rootUrl, "with error:", err.Error())
	}

	c.Wait()
}
