package main

import (
	"chpunk/commands/file"
	"chpunk/commands/sheet"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{Use: "translate"}
	rootCmd.AddCommand(file.Command())
	rootCmd.AddCommand(sheet.Command())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
