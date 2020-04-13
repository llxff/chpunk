package sheets

import (
	"chpunk/internal/export/googledoc"
	"chpunk/internal/google/doc"
	"chpunk/internal/google/files"
	"chpunk/internal/import/sheets"
	"chpunk/internal/settings"
	"chpunk/internal/translation"
	"chpunk/web/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

type indexRequest struct {
	Filter string `json:"filter"`
}

func Index(ctx echo.Context) error {
	var params indexRequest
	if err := ctx.Bind(&params); err != nil {
		return err
	}

	s := driveClient(ctx)

	f, err := s.Spreadsheets(100, params.Filter)
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

	s := driveClient(ctx)

	f, err := s.CreateSpreadsheet(params.Name)
	if err != nil {
		return err
	}

	return ctx.JSON(200, f)
}

func Translate(ctx echo.Context) error {
	sheetID := ctx.Param("sheetID")
	docID := ctx.Param("docID")

	config := settings.Get()
	lines := sheets.Import(sheetID)
	translations := translation.Translate(*config, lines)

	exporter := &googledoc.Container{
		Client: &doc.Client{HTTPClient: googleClient(ctx)},
		DocID:  docID,
	}

	err := exporter.Export(translations)
	if err != nil {
		return err
	}

	return ctx.JSON(200, "OK")
}

func driveClient(ctx echo.Context) *files.Client {
	return &files.Client{HTTPClient: googleClient(ctx)}
}

func googleClient(ctx echo.Context) *http.Client {
	return ctx.Get(middlewares.GoogleClient).(*http.Client)
}
