package file

import (
	"chpunk/export/textfile"
	"chpunk/import/csvfile"
	"chpunk/translation"
	"github.com/spf13/cobra"
	"strings"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "file [path to file] [path to file with translation (default translation.txt)]",
		Short: "Translates a given file",
		Long:  "Translates a given file with Deepl and Yandex translators and saves output to a given file",
		Args:  cobra.MinimumNArgs(1),
		Run:   run,
	}
}

func run(_ *cobra.Command, args []string) {
	lines := csvfile.Import(args[0])
	translations := translation.Translate(lines)
	outputFile := args[1]

	if strings.TrimSpace(outputFile) == "" {
		outputFile = "translation.txt"
	}

	textfile.Export(translations, outputFile)
}