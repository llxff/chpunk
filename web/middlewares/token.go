package middlewares

import (
	"chpunk/google/client"
	"encoding/base64"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"net/http"
)

const GoogleClient = "GoogleClient"

func Token() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			c, err := getGoogleClient(ctx)
			if err != nil {
				return err
			}

			ctx.Set(GoogleClient, c)

			return next(ctx)
		}
	}
}

func getGoogleClient(ctx echo.Context) (*http.Client, error) {
	t, err := getOAuthToken(ctx)
	if err != nil {
		return nil, err
	}

	c, err := client.GetFromToken(t)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized token").SetInternal(err)
	}

	return c, nil
}

func getOAuthToken(ctx echo.Context) (*oauth2.Token, error) {
	tokenB64 := ctx.Request().Header.Get(echo.HeaderAuthorization)
	if tokenB64 == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "No token")
	}

	payload, err := base64.StdEncoding.DecodeString(tokenB64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Wrong token value").SetInternal(err)
	}

	var t oauth2.Token

	err = json.Unmarshal(payload, &t)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token value").SetInternal(err)
	}

	return &t, nil
}
