package doc

import (
	"chpunk/google/client"
	"chpunk/google/docs"
	"github.com/spf13/cobra"
	"log"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "doc [doc id] [path to file with translation (default translation.txt)]",
		Short: "Translates a given Google Spreadsheet",
		Long:  "Translates a given Google Spreadsheet with Deepl and Yandex translators and saves output to a given file",
		Args:  cobra.MinimumNArgs(1),
		Run:   run,
	}
}

func run(_ *cobra.Command, args []string) {
	c := client.Get("token.json")
	d := &docs.Client{HTTPClient: c}

	err := d.NewChapter(args[0], []string{"1", "***", "3"})
	if err != nil {
		log.Fatalln(err)
	}
}
