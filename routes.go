package main

import (
	"net/http"
	"streamcat-api/configuration"
	h "streamcat-api/handlers"
	"streamcat-api/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func registerRoutes(e *echo.Echo) {
	config := configuration.ConfigurationSetup()

	e.GET("/", index)
	e.GET("/stream/auth", h.AuthenticateStream) // RTMP auth check.
	e.POST("/authorize", authorize)             // JWT login.
	e.POST("/register", h.RegisterClient)

	// User routes
	v1 := e.Group("/v1")
	v1.GET("/users", h.GetUsers)
	v1.GET("/users/:id", h.GetUser)
	v1.POST("/users", h.CreateUser)
	v1.PUT("/users/:id", h.UpdateUser)
	v1.DELETE("/users/:id", h.DeleteUser)

	// Stream routes
	v1.GET("/streams", h.GetStreams)
	v1.GET("/streams/stats", h.GetAllStreamStats)
	v1.GET("/streams/featured", h.GetFeaturedStreams)
	v1.GET("/streams/:name", h.GetStreamByName)
	v1.GET("/streams/:name/active", h.StreamActive)
	v1.POST("/streams", h.CreateStream)
	v1.PUT("/streams/:name", h.UpdateStream)
	v1.DELETE("/streams/:name", h.DeleteStream)
	v1.GET("/stream/auth", h.AuthenticateStream)

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
