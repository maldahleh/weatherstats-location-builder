package main

import (
	"log"
	"os"

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

	setupLogging()
	scraper.Scrape("ON")
}

func setupLogging() {
	file, err := os.OpenFile("info.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("test")
}
