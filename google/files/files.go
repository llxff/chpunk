package files

import (
	"fmt"
	"google.golang.org/api/drive/v3"
	"net/http"
)

const (
	SpreadsheetsMimeType = "application/vnd.google-apps.spreadsheet"
	DocumentsMimeType    = "application/vnd.google-apps.document"
)

type Client struct {
	HTTPClient *http.Client
}

func (c *Client) CreateDocument(name string) (*drive.File, error) {
	return c.create(name, DocumentsMimeType)
}

func (c *Client) CreateSpreadsheet(name string) (*drive.File, error) {
	return c.create(name, SpreadsheetsMimeType)
}

func (c *Client) Spreadsheets(pageSize int64, name string) (*drive.FileList, error) {
	return c.files(pageSize, name, SpreadsheetsMimeType)
}

func (c *Client) Documents(pageSize int64, name string) (*drive.FileList, error) {
	return c.files(pageSize, name, DocumentsMimeType)
}

func (c *Client) files(pageSize int64, name string, mimeType string) (*drive.FileList, error) {
	srv, err := drive.New(c.HTTPClient)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	return srv.
		Files.
		List().
		OrderBy("modifiedByMeTime desc").
		Q("trashed=false and mimeType='" + mimeType + "' and name contains '" + name + "'").
		PageSize(pageSize).
		Fields("nextPageToken, files(id, name, mimeType)").
		Do()
}

func (c *Client) create(name string, mimeType string) (*drive.File, error) {
	srv, err := drive.New(c.HTTPClient)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	f := &drive.File{MimeType: mimeType, Name: name}
	return srv.Files.Create(f).Do()
}
