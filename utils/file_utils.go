package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func DeleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Error("failed to delete file", path, "error", err)
	}
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

