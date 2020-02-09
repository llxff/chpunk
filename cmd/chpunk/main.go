package main

import (
	"chpunk/commands/file"
	"chpunk/commands/sheets"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{Use: "translate"}
	rootCmd.AddCommand(file.Command())
	rootCmd.AddCommand(sheets.Command())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
