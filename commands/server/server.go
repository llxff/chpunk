package server

import (
	"chpunk/web/controllers/login"
	"chpunk/web/controllers/sheets"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func run(_ *cobra.Command, args []string) {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
	}))

	e.POST("/auth", login.Auth)
	e.POST("/oauth/callback", login.Callback)
	e.POST("/sheets/:id", sheets.Handle)

	e.Logger.Fatal(e.Start(":" + port))
}
