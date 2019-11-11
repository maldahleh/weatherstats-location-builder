package downloader

import (
	"testing"

	utils "github.com/maldahleh/fileutils"
)

func TestDownloadFile(t *testing.T) {
	const url = "https://dd.weather.gc.ca/climate/observations/daily/csv/NU/climate_daily_NU_2300426_2008-07_PID.csv"
	const path = "test.csv"

	err := DownloadFile(path, url)
	if err != nil {
		t.Fatalf("[TestDownloadFile] file downloading failed with error: %s", err.Error())
	}

	if !utils.FileExists(path) {
		t.Fatalf("[TestDownloadFile] file %s does not exist", path)
	}

	utils.DeleteFile(path)
}

func TestDownloadFileFailure(t *testing.T) {
	const url = "https://dd.weathe.gc.ca/climate/observations/daily/csv/NU/climate_daily_NU_2300426_2008-07_PID.csv"
	const path = "test.csv"

	err := DownloadFile(path, url)
	if err == nil {
		utils.DeleteFile(path)
		t.Fatal("[TestDownloadFileFailure] file downloading did not fail when expected")
	}

	if utils.FileExists(path) {
		utils.DeleteFile(path)
		t.Fatalf("[TestDownloadFileFailure] file %s exists when it should not", path)
	}
}
