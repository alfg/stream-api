package main

import (
	"fmt"
	"runtime"
	"stream-api/configuration"

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
	config := configuration.ConfigurationSetup()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// Routes
	registerRoutes(e)

	// Start server
	fmt.Printf("Starting server on port %s\n", config.Port)
	// e.Run(standard.New(fmt.Sprintf("%s:%s", config.Host, config.Port)))
	// e.Run(standard.New(":4000"))
	e.Logger.Fatal(e.Start(":4000"))
}

func main() {
	configRuntime()
	startServer()
}
