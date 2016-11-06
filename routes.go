package main

import (
	"net/http"
	"stream-api/controllers"

	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) {
	// config := ConfigurationSetup()

	e.Get("/", index)
	e.Get("/stream/auth", auth)

	// User routes
	v1 := e.Group("/v1")
	v1.Get("/users", controllers.GetUsers)
	v1.Get("/users/:id", controllers.GetUser)
	v1.Post("/users", controllers.CreateUser)
	v1.Put("/users/:id", controllers.UpdateUser)
	v1.Delete("/users/:id", controllers.DeleteUser)

	// Stream routes
	v1.Get("/stream", controllers.GetStreams)
	v1.Get("/stream/:id", controllers.GetStream)
	v1.Post("/stream", controllers.CreateStream)
	v1.Put("/stream/:id", controllers.UpdateStream)
	v1.Delete("/stream/:id", controllers.DeleteStream)

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

func auth(c echo.Context) error {
	key := c.QueryParam("key")
	if key == "testkey" {
		return c.String(http.StatusOK, "OK")
	}
	return c.String(http.StatusForbidden, "Forbidden")
}
