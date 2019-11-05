package downloader

import (
	"testing"

	"weatherstatsLocations/utils"
)

func TestDownloadFile(t *testing.T) {
	const testUrl = "https://dd.weather.gc.ca/climate/observations/daily/csv/NU/climate_daily_NU_2300426_2008-07_PID.csv"
	const outPath = "test.csv"

	err := DownloadFile(outPath, testUrl)
	if err != nil {
		t.Fatalf("[TestDownloadFile] file downloading failed with error: %s", err.Error())
	}

	if !utils.FileExists(outPath) {
		t.Fatalf("[TestDownloadFile] file %s does not exist", outPath)
	}

	utils.DeleteFile(outPath)
}

func TestDownloadFileFailure(t *testing.T) {
	const testUrl = "https://dd.weathe.gc.ca/climate/observations/daily/csv/NU/climate_daily_NU_2300426_2008-07_PID.csv"
	const outPath = "test.csv"

	err := DownloadFile(outPath, testUrl)
	if err == nil {
		t.Fatalf("[TestDownloadFileFailure] file downloading did not fail when expected")
	}

	if utils.FileExists(outPath) {
		utils.DeleteFile(outPath)
		t.Fatalf("[TestDownloadFileFailure] file %s exists when it should not", outPath)
	}
}
