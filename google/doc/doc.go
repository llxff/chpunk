package doc

import (
	"google.golang.org/api/docs/v1"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
}

func (c *Client) Service() (*docs.Service, error) {
	return docs.New(c.HTTPClient)
}
