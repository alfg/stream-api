package main

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
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
	config := ConfigurationSetup()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// Routes
	registerRoutes(e)

	// Start server
	fmt.Printf("Starting server on port %s\n", config.Port)
	// e.Run(standard.New(fmt.Sprintf("%s:%s", config.Host, config.Port)))
	e.Run(standard.New(":4000"))
}

func main() {
	configRuntime()
	startServer()
}
