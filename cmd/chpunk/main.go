package main

import (
	"chpunk/cmd/doc"
	"chpunk/cmd/file"
	"chpunk/cmd/server"
	"chpunk/cmd/sheet"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "translate"}
	rootCmd.AddCommand(file.Command())
	rootCmd.AddCommand(sheet.Command())
	rootCmd.AddCommand(server.Command())
	rootCmd.AddCommand(doc.Command())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
