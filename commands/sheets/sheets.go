package sheets

import (
	"chpunk/google/client"
	"chpunk/google/files"
	"chpunk/google/spreadsheets"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "sheets",
		Short: "List of available Google Sheets",
		Long:  "List of available Google Sheets",
		Args:  cobra.NoArgs,
		Run:   run,
	}
}

func run(_ *cobra.Command, _ []string) {
	c := client.Get("token.json")

	d := &files.Client{HttpClient: c}
	s := &spreadsheets.Client{HttpClient: c}

	r, err := d.Files()

	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		data := s.Values(r.Files[0].Id, "A2:B")

		if len(data) == 0 {
			fmt.Println("No data found.")
		} else {
			fmt.Println("A, B:")
			for _, row := range data {
				fmt.Printf("%s, %s\n", row[0], row[1])
			}
		}

		fmt.Println("Files:")
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.MimeType)
		}
	}
}
