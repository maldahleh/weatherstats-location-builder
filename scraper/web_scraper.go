package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

const rootUrl string = "http://dd.weatheroffice.ec.gc.ca/climate/observations/daily/csv/"

func Scrape() {
	c := colly.NewCollector(
		colly.AllowedDomains(
			"dd.weatheroffice.ec.gc.ca",
			"www.dd.weatheroffice.ec.gc.ca",
		),
		colly.CacheDir("./eccc_cache"),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(rootUrl)
	if err != nil {
		fmt.Print("Failed to visit: " + rootUrl + ", error: " + err.Error())
	}
}
