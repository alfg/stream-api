package main

import (
	"net/http"
	"stream-api/services"

	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) {
	// config := ConfigurationSetup()

	e.Get("/", index)

	// User routes
	e.Get("/users", services.GetUsers)
	e.Get("/users/:id", services.GetUser)
	e.Post("/users", services.CreateUser)
	e.Put("/users/:id", services.UpdateUser)
	e.Delete("/users/:id", services.DeleteUser)

	// Stream routes
	e.Get("/stream", services.GetStreams)
	e.Get("/stream/:id", services.GetStream)
	e.Post("/stream", services.CreateStream)
	e.Put("/stream/:id", services.UpdateStream)
	e.Delete("/stream/:id", services.DeleteStream)

	// Restricted group
	// Temporary: Run scripts/token.go to generate auth token
	r := e.Group("/restricted")
	// r.Use(JWTAuth(config.JWTKey))
	r.Get("", restricted)
}

// Handlers
func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
