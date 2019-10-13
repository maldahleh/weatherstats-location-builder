package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
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
		colly.CacheDir("./eccc_cache"),
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

		fmt.Println(split)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(rootUrl + province + "/")
	if err != nil {
		fmt.Println("Failed to visit:", rootUrl, "with error:", err.Error())
	}
}
