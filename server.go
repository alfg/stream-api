package main

import (
	"fmt"
	"runtime"
	"streamcat-api/settings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func configRuntime() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Running with %d CPUs\n", numCPU)
}

func startServer() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// Routes
	registerRoutes(e)

	// Start server
	port := fmt.Sprintf(":%s", settings.Get().Port)
	e.Logger.Fatal(e.Start(port))
}

func main() {
	configRuntime()
	startServer()
}
