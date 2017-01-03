package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"streamcat-api/configuration"
	"streamcat-api/data"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func authorize(c echo.Context) error {
	config := configuration.ConfigurationSetup()
	authorizationHeader := c.Request().Header.Get("Authorization")

	if !strings.HasPrefix(authorizationHeader, "Basic") {
		return echo.ErrUnauthorized
	}

	authorizationHeader = strings.TrimPrefix(authorizationHeader, "Basic ")
	decoded, err := base64.StdEncoding.DecodeString(authorizationHeader)
	if err != nil {
		fmt.Println("error decoding.")
		return echo.ErrUnauthorized
	}

	authArr := strings.Split(string(decoded), ":")

	key := authArr[0]
	secret := authArr[1]

	if data.ValidateClient(key, secret) {

		// Get client from DB.
		client, _ := data.GetClientByKey(key)

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// TODO: Set proper claims for client in DB.
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = client.Email
		claims["domain"] = client.Domain
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(config.JWTKey))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

// func login(c echo.Context) error {
// 	config := configuration.ConfigurationSetup()
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")
//
// 	// Bind model to json request body.
// 	// s := new(models.Stream)
// 	// if err := c.Bind(s); err != nil {
// 	// 	return err
// 	// }
//
// 	if username == "jon" && password == "shhh!" {
// 		// Create token
// 		token := jwt.New(jwt.SigningMethodHS256)
//
// 		// Set claims
// 		claims := token.Claims.(jwt.MapClaims)
// 		claims["name"] = "Jon Snow"
// 		claims["admin"] = true
// 		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
//
// 		// Generate encoded token and send it as response.
// 		t, err := token.SignedString([]byte(config.JWTKey))
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(http.StatusOK, map[string]string{
// 			"token": t,
// 		})
// 	}
//
// 	return echo.ErrUnauthorized
// }

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	admin := claims["admin"].(bool)
	email := claims["email"].(string)
	domain := claims["domain"].(string)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"admin":  admin,
		"email":  email,
		"domain": domain,
	})
}
