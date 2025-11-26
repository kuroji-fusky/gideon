package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// A wrapper for e.GET to include the necessary query params for displaying data and bypassing cache
// func GetReqWrapper(e *echo.Echo, path string, h echo.HandlerFunc) *echo.Route {
// 	return e.GET(path, h)
// }

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {

		// header := c.Response().Header()
		// header.Add("X-Gideon-Cache", "HIT")

		return c.JSON(http.StatusOK, map[string]any{
			"message": "hi",
		})
	})

	// For prefetching links on the server
	e.POST("/prefetch-links", func(c echo.Context) error {})

	// Hives; shown as "Collections" in the UI
	e.GET("/hive/:hive", func(c echo.Context) error {
		noMinify := (c.QueryParam("no_minify"))
	})
	e.POST("/add/hive", func(c echo.Context) error {})
	e.DELETE("/hive/:hive", func(c echo.Context) error {})

	// Page operations
	e.GET("/hive/:hive/:page", func(c echo.Context) error {
		noMinify := (c.QueryParam("no_minify"))
	})
	e.POST("/add/hive/:hive/:page", func(c echo.Context) error {})
	e.DELETE("/hive/:hive/:page", func(c echo.Context) error {})
}
