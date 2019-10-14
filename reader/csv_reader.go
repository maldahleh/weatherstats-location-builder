package reader

import (
	"bufio"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func RetrieveStationName(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.Contains(text, "Station Name") {
			continue
		}

		split := strings.Split(text, ",")
		if len(split) != 2 {
			continue
		}

		return split[1][1 : len(split[1]) - 1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ""
}
