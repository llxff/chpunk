package server

import (
	"chpunk/web/controllers/login"
	"chpunk/web/controllers/sheets"
	"chpunk/web/middlewares"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(port string) {
	e := echo.New()

	setupMiddleware(e)
	setupRoutes(e)

	printRoutes(e.Routes())

	e.Logger.Fatal(e.Start(":" + port))
}

func setupMiddleware(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
	}))
}

func setupRoutes(e *echo.Echo) {
	e.POST("/auth", login.Auth)
	e.POST("/oauth/callback", login.Callback)

	s := e.Group("/sheets")

	s.Use(middlewares.Token())

	s.POST("", sheets.Index)
	s.POST("/:id", sheets.Get)
}

func printRoutes(routes []*echo.Route) {
	for _, route := range routes {
		switch route.Method {
		case "GET", "POST", "PUT", "PATCH", "DELETE":
			fmt.Printf("%-6s %s\n", route.Method, route.Path)
		}
	}
}
