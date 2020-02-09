package files

import (
	"google.golang.org/api/drive/v3"
	"net/http"

	"log"
)

type Client struct {
	HttpClient *http.Client
}

func (c *Client) Files() (*drive.FileList, error) {
	srv, err := drive.New(c.HttpClient)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return srv.
		Files.
		List().
		Q("mimeType='application/vnd.google-apps.spreadsheet'").
		PageSize(10).
		Fields("nextPageToken, files(id, name, mimeType)").
		Do()
}
