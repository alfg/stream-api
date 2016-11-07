package main

import (
	"net/http"
	"stream-api/controllers"
	"stream-api/models"

	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) {
	// config := ConfigurationSetup()

	e.Get("/", index)
	e.Get("/stream/auth", controllers.AuthenticateStream)

	// User routes
	v1 := e.Group("/v1")
	v1.Get("/users", controllers.GetUsers)
	v1.Get("/users/:id", controllers.GetUser)
	v1.Post("/users", controllers.CreateUser)
	v1.Put("/users/:id", controllers.UpdateUser)
	v1.Delete("/users/:id", controllers.DeleteUser)

	// Stream routes
	v1.Get("/streams", controllers.GetStreams)
	v1.Get("/streams/:id", controllers.GetStream)
	v1.Post("/streams", controllers.CreateStream)
	v1.Put("/streams/:id", controllers.UpdateStream)
	v1.Delete("/streams/:id", controllers.DeleteStream)
	v1.Get("/stream/auth", controllers.AuthenticateStream)

	// Restricted group
	// Temporary: Run scripts/token.go to generate auth token
	r := e.Group("/restricted")
	// r.Use(JWTAuth(config.JWTKey))
	r.Get("", restricted)
}

// Handlers
func index(c echo.Context) error {
	i := models.Index{
		Name:    "stream-api",
		Version: "0.0.1",
	}
	return c.JSON(http.StatusOK, &i)
}
