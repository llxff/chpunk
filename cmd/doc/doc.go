package doc

import (
	"chpunk/internal/export/googledoc"
	"chpunk/internal/google/client"
	"chpunk/internal/google/doc"
	"chpunk/internal/import/sheets"
	"chpunk/internal/settings"
	"chpunk/internal/translation"
	"github.com/spf13/cobra"
	"log"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "doc [spreadsheet id] [doc id]",
		Short: "Translates a given Google Spreadsheet and past translation to a given Google Doc",
		Args:  cobra.MinimumNArgs(2),
		Run:   run,
	}
}

func run(_ *cobra.Command, args []string) {
	config := settings.Get()
	lines := sheets.Import(args[0])
	translations := translation.Translate(*config, lines)

	c := client.Get("token.json")
	exporter := &googledoc.Container{
		Client: &doc.Client{HTTPClient: c},
		DocID:  args[1],
	}

	err := exporter.Export(translations)
	if err != nil {
		log.Fatalln(err)
	}
}
