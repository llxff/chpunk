package sheets

import (
	"chpunk/google/files"
	"chpunk/google/spreadsheets"
	"chpunk/web/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(ctx echo.Context) error {
	c := ctx.Get(middlewares.GoogleClient).(*http.Client)
	s := files.Client{HTTPClient: c}

	f, err := s.Files(100)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, f)
}

func Get(ctx echo.Context) error {
	c := ctx.Get(middlewares.GoogleClient).(*http.Client)

	s := &spreadsheets.Client{HTTPClient: c}

	data := s.Values(ctx.Param("id"), "A1:A")

	return ctx.JSON(http.StatusOK, data)
}
