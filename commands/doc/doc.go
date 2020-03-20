package doc

import (
	"chpunk/google/client"
	"chpunk/google/docs"
	"chpunk/import/sheets"
	"chpunk/settings"
	"chpunk/translation"
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
	d := &docs.Client{HTTPClient: c}

	err := d.NewChapter(args[1], translations)
	if err != nil {
		log.Fatalln(err)
	}
}
