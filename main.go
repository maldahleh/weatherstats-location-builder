package main

import (
	"fmt"
	"weatherstatsLocations/scraper"
)

func main() {
	//err := downloader.DownloadFile(
	//	"test.csv",
	//	"http://dd.weatheroffice.ec.gc.ca/climate/observations/daily/csv/NU/climate_daily_NU_2403945_2015-05_PID.csv",
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}

	m := scraper.Scrape("ON")
	for k, v := range m {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
}
