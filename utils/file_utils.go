package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func DeleteFile(filepath string) {
	err := os.Remove(filepath)
	if err != nil {
		log.Error("failed to delete file", filepath, "error", err)
	}
}

func FileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

