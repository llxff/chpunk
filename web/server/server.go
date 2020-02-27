package server

import (
	"chpunk/web/controllers/login"
	"chpunk/web/controllers/sheets"
	"chpunk/web/middlewares"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func Start(port string) {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
	}))

	e.POST("/auth", login.Auth)
	e.POST("/oauth/callback", login.Callback)

	s := e.Group("/sheets")

	s.Use(middlewares.Token())

	s.POST("", sheets.Index)
	s.POST("/:id", sheets.Get)

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	e.Logger.Fatal(e.Start(":" + port))
}
