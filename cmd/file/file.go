package file

import (
	"chpunk/internal/export/textfile"
	"chpunk/internal/import/csvfile"
	"chpunk/internal/settings"
	"chpunk/internal/translation"
	"github.com/spf13/cobra"
	"log"
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
	config := settings.Get()
	lines := csvfile.Import(args[0])
	translations := translation.Translate(*config, lines)
	outputFile := args[1]

	if strings.TrimSpace(outputFile) == "" {
		outputFile = "translation.txt"
	}

	exporter := &textfile.Container{FileName: outputFile}
	err := exporter.Export(translations)

	if err != nil {
		log.Fatalln(err)
	}
}
