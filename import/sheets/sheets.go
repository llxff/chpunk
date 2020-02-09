package sheets

import (
	"chpunk/google/client"
	"chpunk/google/spreadsheets"
	"fmt"
)

func Import(fileName string) (lines []string) {
	c := client.Get("token.json")
	s := &spreadsheets.Client{HttpClient: c}

	data := s.Values(fileName, "A1:A")

	if len(data) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, l := range data {
			lines = append(lines, l[0].(string))
		}
	}

	return
}