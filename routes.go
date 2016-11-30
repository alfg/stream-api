package main

import (
	"net/http"
	"stream-api/controllers"
	"stream-api/models"

	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) {
	// config := ConfigurationSetup()

	e.GET("/", index)
	e.GET("/stream/auth", controllers.AuthenticateStream)

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
	// Temporary: Run scripts/token.go to generate auth token
	r := e.Group("/restricted")
	// r.Use(JWTAuth(config.JWTKey))
	r.GET("", restricted)
}

// Handlers
func index(c echo.Context) error {
	i := models.Index{
		Name:    "stream-api",
		Version: "0.0.1",
	}
	return c.JSON(http.StatusOK, &i)
}
