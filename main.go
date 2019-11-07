package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"weatherstatsLocations/scraper"

	log "github.com/sirupsen/logrus"
)

var mu sync.RWMutex
var resp []byte

func allStations(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	b := resp
	mu.RUnlock()

	_, err := w.Write(b)
	if err != nil {
		log.Error("HTTP Write Failure", err)
	}
}

func handleRequests() {
	http.HandleFunc("/stations", allStations)

	err := http.ListenAndServe(":8080", nil)
	log.Fatal("Encountered error with web server", err)
}

func main() {
	provinceData := scraper.Scrape()
	resp, _ = json.Marshal(provinceData)

	go func() {
		for {
			time.Sleep(time.Second * 1000)
			mu.Lock()
			provinceData = scraper.Scrape()
			resp, _ = json.Marshal(provinceData)
			mu.Unlock()
		}
	}()

	handleRequests()
}
