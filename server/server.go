package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getEnvWithFallback[T any](envKey string, fallback T) T {
	key := os.Getenv(envKey)

	if key == "" {
		return fallback
	}

	switch any(fallback).(type) {
	case string:
		return any(key).(T)
	default:
		return fallback
	}
}

type DomainConfig struct {
	AllowedDomains     []string  `yaml:"allowed_domains"`
	BlacklistedDomains *[]string `yaml:"blacklisted_domains"`
}

func main() {
	// Load env
	cwd, _ := os.Getwd()

	envErr := godotenv.Load(filepath.Join(cwd, ".env"))
	if envErr != nil {
		log.Fatalf("Error getting .env: %v", envErr)
	}

	customAPIToken := os.Getenv("GIDEON_TOKEN")
	serverPort := getEnvWithFallback("GIDEON_SERVER_PORT", 6969)
	allowedURLOrigins := getEnvWithFallback("GIDEON_SERVER_CORS_ALLOWED_DOMAINS", "http://localhost:4321")

	fmt.Println("Token:", customAPIToken)
	fmt.Println("Server port:", serverPort)

	e := echo.New()

	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     strings.Split(allowedURLOrigins, ","),
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true,
		}),
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: 20 * time.Second,
		}),
	)

	// Parse domains.yml
	cfgBuffer, cfgFileErr := os.ReadFile(filepath.Join(cwd, "domains.yml"))
	if cfgFileErr != nil {
		log.Panicf("There was an oopsie reading domains.yml", cfgFileErr)
	}

	var loadedDomainCfg DomainConfig

	domainCfgLoadErr := yaml.Unmarshal(cfgBuffer, loadedDomainCfg)
	if domainCfgLoadErr != nil {
		log.Panicf("Cannot load domain config", domainCfgLoadErr)
	}

	// Load dem routes
	Routes(e)

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", serverPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	closeDatMf := make(chan os.Signal, 1)
	signal.Notify(closeDatMf, os.Interrupt)

	<-closeDatMf

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("Server forced to shut down: %v", err)
	}
}
