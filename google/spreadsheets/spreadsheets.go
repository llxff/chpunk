package spreadsheets

import (
	"log"
	"net/http"

	"google.golang.org/api/sheets/v4"
)

type Client struct {
	HTTPClient *http.Client
}

func (c *Client) Values(spreadsheetID string, readRange string) [][]interface{} {
	srv, err := sheets.New(c.HTTPClient)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	return resp.Values
}
