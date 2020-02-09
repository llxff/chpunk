package spreadsheets

import (
	"google.golang.org/api/sheets/v4"
	"log"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
}

func (c *Client) Values(spreadsheetId string, readRange string) [][]interface{} {
	srv, err := sheets.New(c.HttpClient)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	return resp.Values
}
