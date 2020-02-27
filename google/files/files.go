package files

import (
	"google.golang.org/api/drive/v3"
	"net/http"

	"log"
)

type Client struct {
	HTTPClient *http.Client
}

func (c *Client) Files(pageSize int64) (*drive.FileList, error) {
	srv, err := drive.New(c.HTTPClient)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return srv.
		Files.
		List().
		OrderBy("modifiedByMeTime desc").
		Q("mimeType='application/vnd.google-apps.spreadsheet'").
		PageSize(pageSize).
		Fields("nextPageToken, files(id, name, mimeType)").
		Do()
}
