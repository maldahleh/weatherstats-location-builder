package main

import (
	"fmt"

	"weatherstatsLocations/scraper"
)

func main() {
	m := scraper.Scrape("ON")
	for k, v := range m {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
}
