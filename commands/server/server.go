package server

import (
	"chpunk/web/server"
	"github.com/spf13/cobra"
)

var port string

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Starts server",
		Long:  "Starts API server",
		Args:  cobra.NoArgs,
		Run:   run,
	}

	cmd.Flags().StringVarP(&port, "port", "p", "80", "Port")

	return cmd
}

func run(_ *cobra.Command, _ []string) {
	server.Start(port)
}
