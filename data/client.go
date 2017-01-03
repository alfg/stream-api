package data

import (
	"fmt"
	"streamcat-api/models"

	"golang.org/x/crypto/bcrypt"
)

func ValidateClient(key string, secret string) bool {
	// err = bcrypt.CompareHashAndPassword(secret, password)
	client, err := GetClientByKey(key)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(client.APISecret), []byte(secret))
	if err == nil {
		fmt.Println("Authorized.")
		return true
	}
	return false
}

// GetClientByKey Gets client by key.
func GetClientByKey(key string) (*models.Client, error) {
	const query = `SELECT * FROM client WHERE api_key = $1`

	db, _ := ConnectDB()
	client := models.Client{}
	err := db.Get(&client, query, key)
	fmt.Println(&client)
	if err != nil {
		fmt.Println(err)
		return &client, err
	}

	return &client, nil
}

// CreateClient creates a client.
func CreateClient(client *models.Client) (*models.Client, error) {
	const query = "INSERT INTO client (api_key, api_secret, email, domain) VALUES (:api_key, :api_secret, :email, :domain)"

	// Create bcrypt hashed password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(client.APISecret), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	client.APISecret = string(hashedPassword)

	db, _ := ConnectDB()
	tx := db.MustBegin()
	result, err := tx.NamedExec(query, &client)
	if err != nil {
		fmt.Println("Error", err)
		return client, err
	}
	tx.Commit()

	fmt.Println("transaction done")

	lastID, _ := result.LastInsertId()
	client.ID = lastID

	return client, nil
}

// ClientExistsByEmail Checks if email exists.
func ClientExistsByEmail(email string) bool {
	const query = `SELECT EXISTS (SELECT id FROM client WHERE email = $1)`

	var exists bool
	db, _ := ConnectDB()
	err := db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		fmt.Println(err)
	}
	return exists
}
