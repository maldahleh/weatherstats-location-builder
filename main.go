package main

import (
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

	scraper.Scrape("ON")
}
