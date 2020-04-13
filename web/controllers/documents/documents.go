package documents

import (
	"chpunk/internal/google/files"
	"chpunk/web/middlewares"
	"github.com/labstack/echo/v4"
	"net/http"
)

type indexRequest struct {
	Filter string `json:"filter"`
}

func Index(ctx echo.Context) error {
	var params indexRequest
	if err := ctx.Bind(&params); err != nil {
		return err
	}

	s := googleClient(ctx)

	f, err := s.Documents(100, params.Filter)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, f)
}

type createRequest struct {
	Name string `json:"name"`
}

func Create(ctx echo.Context) error {
	var params createRequest
	if err := ctx.Bind(&params); err != nil {
		return err
	}

	s := googleClient(ctx)

	f, err := s.CreateDocument(params.Name)
	if err != nil {
		return err
	}

	return ctx.JSON(200, f)
}

func googleClient(ctx echo.Context) *files.Client {
	c := ctx.Get(middlewares.GoogleClient).(*http.Client)
	return &files.Client{HTTPClient: c}
}
