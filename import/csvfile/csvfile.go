package csvfile

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func Import(fileName string) []string {
	csvfile, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	lines := make([]string, 0, 10)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		line := strings.TrimSpace(record[0])

		if line != "" {
			lines = append(lines, line)
		}
	}

	return lines
}
