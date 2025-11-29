package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// A wrapper for e.GET to include the necessary query params for displaying data and bypassing cache
// func GetReqWrapper(e *echo.Echo, path string, h echo.HandlerFunc) *echo.Route {
// 	return e.GET(path, h)
// }

// Parse auth token
func parseAuthHeader(ec echo.Context) (string, error) {
	authHeader := ec.Request().Header.Get("Authorization")

	if authHeader == "" {
		return "", ec.JSON(http.StatusUnauthorized, map[string]any{
			"message": "Auth header required",
		})
	}

	splitAuth := strings.Split(authHeader, "Bearer ")

	return splitAuth[1], nil
}

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		// header := c.Response().Header()
		// header.Add("X-Gideon-Cache", "HIT")

		return c.JSON(http.StatusOK, map[string]any{
			"message": "hi",
		})
	})

	// For prefetching links on the server
	e.GET("/prefetch-links", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"links": []string{"dummy url"},
		})
	})

	e.POST("/prefetch-links", func(c echo.Context) error {
		authToken, _ := parseAuthHeader(c)

		postMsg := make(map[string]struct {
			Links string `json:"links"`
		})

		json_err := json.NewDecoder(c.Request().Body).Decode(&postMsg)
		if json_err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		return c.NoContent(http.StatusOK)
	})

	// Hives; shown as "Collections" in the UI
	e.GET("/hive/:hive", func(c echo.Context) error {
		noMinify, _ := strconv.Atoi(c.QueryParam("no_minify"))
	})

	e.POST("/add/hive", func(c echo.Context) error {
		authToken, _ := parseAuthHeader(c)
	})

	e.PATCH("/update/:hive", func(c echo.Context) error {
		authToken, _ := parseAuthHeader(c)
	})

	e.DELETE("/hive/:hive", func(c echo.Context) error {
		authToken, _ := parseAuthHeader(c)
	})

	// Page operations
	e.GET("/hive/:hive/:page", func(c echo.Context) error {
		noMinify, _ := strconv.Atoi(c.QueryParam("no_minify"))

		hiveId := c.Param("hive")
		pageSlug := c.Param("page")
	})

	e.POST("/hive/:hive/:page", func(c echo.Context) error {
		authToken, _ := parseAuthHeader(c)

		hiveId := c.Param("hive")
		pageSlug := c.Param("page")
	})

	e.PATCH("/hive/:hive/:page", func(c echo.Context) error {
		authToken, _ := parseAuthHeader(c)

		hiveId := c.Param("hive")
		pageSlug := c.Param("page")
	})

	e.DELETE("/hive/:hive/:page", func(c echo.Context) error {
		authToken, _ := parseAuthHeader(c)
		isPermDel, _ := strconv.Atoi(c.QueryParam("perm"))

		hiveId := c.Param("hive")
		pageSlug := c.Param("page")
	})
}
