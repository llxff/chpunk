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
		Short: "Translate a given file",
		Long:  "Translate a given file with Deepl and Yandex translators and save output to file",
		Args:  cobra.MinimumNArgs(1),
		Run:   run,
	}
}

func run(_ *cobra.Command, args []string) {
	lines := csvfile.Import(args[1])
	translations := translation.Translate(lines)
	outputFile := args[2]

	if strings.TrimSpace(outputFile) == "" {
		outputFile = "translation.txt"
	}

	textfile.Export(translations, outputFile)
}
