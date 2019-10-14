package main

import (
	"fmt"

	"weatherstatsLocations/scraper"
)

func main() {
	m := scraper.ScrapeProvinces()
	for k, v := range m {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
}
