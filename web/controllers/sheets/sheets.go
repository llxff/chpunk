package sheets

import (
	"chpunk/google/client"
	"chpunk/google/spreadsheets"
	"encoding/base64"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"net/http"
)

type req struct {
	Token string `json:"token"`
}

func (r *req) oauthToken() (*oauth2.Token, error) {
	payload, err := base64.StdEncoding.DecodeString(r.Token)
	if err != nil {
		return nil, err
	}

	var t oauth2.Token
	err = json.Unmarshal(payload, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func Handle(c echo.Context) error {
	r := new(req)
	if err := c.Bind(r); err != nil {
		return err
	}

	token, err := r.oauthToken()
	if err != nil {
		return err
	}

	conf, err := client.GetFromToken(token)
	if err != nil {
		return err
	}

	s := &spreadsheets.Client{HttpClient: conf}
	data := s.Values(c.Param("id"), "A1:A")

	return c.JSON(http.StatusOK, data)
}
