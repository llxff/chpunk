package sheet

import (
	"chpunk/export/textfile"
	"chpunk/import/sheets"
	"chpunk/settings"
	"chpunk/translation"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "sheet [doc id] [path to file with translation (default translation.txt)]",
		Short: "Translates a given Google Spreadsheet",
		Long:  "Translates a given Google Spreadsheet with Deepl and Yandex translators and saves output to a given file",
		Args:  cobra.MinimumNArgs(1),
		Run:   run,
	}
}

func run(_ *cobra.Command, args []string) {
	config := settings.Get()
	lines := sheets.Import(args[0])
	translations := translation.Translate(*config, lines)
	outputFile := "translation.txt"

	if len(args) > 1 {
		outputFile = args[1]
	}

	textfile.Export(translations, outputFile)
}
