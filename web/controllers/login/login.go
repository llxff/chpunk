package login

import (
	"chpunk/internal/google/auth"
	"encoding/base64"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Auth(c echo.Context) error {
	return c.String(http.StatusOK, auth.GetAuthURL())
}

type token struct {
	Code string `json:"code"`
}

func Callback(c echo.Context) error {
	code := new(token)
	if err := c.Bind(code); err != nil {
		return err
	}

	oauthToken, err := auth.Authenticate(code.Code)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(oauthToken)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, base64.StdEncoding.EncodeToString(bytes))
}
