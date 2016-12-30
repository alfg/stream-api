package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"streamcat-api/data"
	"streamcat-api/models"

	"github.com/labstack/echo"
)

// GetUser Gets user.
// func GetClient(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
//
// 	user, err := data.GetUserByID(id)
// 	if err != nil {
// 		fmt.Println(err)
// 		resp := models.DoesNotExist{
// 			Code:    404,
// 			Message: "User does not exist",
// 		}
// 		return c.JSON(http.StatusNotFound, resp)
// 	}
// 	fmt.Println(user)
//
// 	return c.JSON(http.StatusOK, user)
// }

// RegisterClient Creates a client
func RegisterClient(c echo.Context) error {

	key := generateKey(10)
	secret := generateKey(20)

	cl := models.Client{
		APIKey:    key,
		APISecret: secret,
	}

	// Create record
	client, err := data.CreateClient(cl)
	if err != nil {
		fmt.Println(client)
	}
	return c.JSON(http.StatusCreated, cl)
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
