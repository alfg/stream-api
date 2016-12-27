package main

import (
	"fmt"
	"net/http"
	"streamcat-api/configuration"
	"streamcat-api/controllers"
	"streamcat-api/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func registerRoutes(e *echo.Echo) {
	config := configuration.ConfigurationSetup()

	e.GET("/", index)
	e.GET("/stream/auth", controllers.AuthenticateStream) // RTMP auth check.
	e.POST("/login", login)                               // JWT login.

	// User routes
	v1 := e.Group("/v1")
	v1.GET("/users", controllers.GetUsers)
	v1.GET("/users/:id", controllers.GetUser)
	v1.POST("/users", controllers.CreateUser)
	v1.PUT("/users/:id", controllers.UpdateUser)
	v1.DELETE("/users/:id", controllers.DeleteUser)

	// Stream routes
	v1.GET("/streams", controllers.GetStreams)
	v1.GET("/streams/stats", controllers.GetAllStreamStats)
	// v1.Get("/streams/:id", controllers.GetStream)
	v1.GET("/streams/featured", controllers.GetFeaturedStreams)
	v1.GET("/streams/:name", controllers.GetStreamByName)
	v1.GET("/streams/:name/active", controllers.StreamActive)
	v1.POST("/streams", controllers.CreateStream)
	v1.PUT("/streams/:id", controllers.UpdateStream)
	v1.DELETE("/streams/:id", controllers.DeleteStream)
	v1.GET("/stream/auth", controllers.AuthenticateStream)

	// Restricted group
	r := e.Group("/restricted")
	fmt.Println(config.JWTKey)
	r.Use(middleware.JWT([]byte(config.JWTKey)))
	r.GET("", restricted)
}

// Handlers
func index(c echo.Context) error {
	i := models.Index{
		Name:    "streamcat-api",
		Version: "0.0.1",
	}
	return c.JSON(http.StatusOK, &i)
}
