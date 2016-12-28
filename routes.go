package main

import (
	"net/http"
	"streamcat-api/configuration"
	"streamcat-api/handlers"
	"streamcat-api/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func registerRoutes(e *echo.Echo) {
	config := configuration.ConfigurationSetup()

	e.GET("/", index)
	e.GET("/stream/auth", handlers.AuthenticateStream) // RTMP auth check.
	e.POST("/login", login)                            // JWT login.

	// User routes
	v1 := e.Group("/v1")
	v1.GET("/users", handlers.GetUsers)
	v1.GET("/users/:id", handlers.GetUser)
	v1.POST("/users", handlers.CreateUser)
	v1.PUT("/users/:id", handlers.UpdateUser)
	v1.DELETE("/users/:id", handlers.DeleteUser)

	// Stream routes
	v1.GET("/streams", handlers.GetStreams)
	v1.GET("/streams/stats", handlers.GetAllStreamStats)
	v1.GET("/streams/featured", handlers.GetFeaturedStreams)
	v1.GET("/streams/:name", handlers.GetStreamByName)
	v1.GET("/streams/:name/active", handlers.StreamActive)
	v1.POST("/streams", handlers.CreateStream)
	v1.PUT("/streams/:name", handlers.UpdateStream)
	v1.DELETE("/streams/:name", handlers.DeleteStream)
	v1.GET("/stream/auth", handlers.AuthenticateStream)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte(config.JWTKey)))
	r.GET("", restricted)
}

// Handlers
func index(c echo.Context) error {
	i := models.Index{
		Name:    "streamcat-api",
		Version: "0.0.1",
		Docs:    "http://docs.streamcat.tv/",
		Github:  "https://github.com/streamcatTV",
	}
	return c.JSON(http.StatusOK, &i)
}
