package data

import (
	"fmt"
	"stream-api/models"
)

// User model

// GetUserByID Gets user by Id
func GetUserByID(id int) (*models.User, error) {
	const query = `SELECT * FROM user WHERE id = $1`

	db, _ := ConnectDB()
	user := models.User{}
	err := db.Get(&user, query, id)
	fmt.Println(&user)
	if err != nil {
		fmt.Println(err)
		return &user, err
	}

	return &user, nil
}

// GetUsers Gets all users.
func GetUsers() *[]models.User {
	const query = `SELECT * FROM user ORDER BY id ASC`

	db, _ := ConnectDB()
	user := []models.User{}
	db.Select(&user, query)
	fmt.Println(&user)

	return &user
}

// CreateUser creates a user.
func CreateUser(user models.User) *models.User {
	const query = "INSERT INTO user (first_name, last_name, email) VALUES (:first_name, :last_name, :email)"

	db, _ := ConnectDB()
	tx := db.MustBegin()
	result, err := tx.NamedExec(query, &user)
	if err != nil {
		fmt.Println("Error", err)
	}
	tx.Commit()

	fmt.Println("transaction done")

	lastID, _ := result.LastInsertId()
	user.ID = lastID

	return &user
}

// UpdateUserByID Update user by id
func UpdateUserByID(id int, user models.User) *models.User {
	const query = `UPDATE user
		SET first_name = :first_name,
		last_name = :last_name,
		email = :email
		WHERE id = :id`

	db, _ := ConnectDB()
	tx := db.MustBegin()
	_, err := tx.NamedExec(query, &user)
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()

	return &user
}

// DeleteUserByID Deletes user by id
func DeleteUserByID(id int) error {
	const query = "DELETE FROM user WHERE id = :id"

	db, _ := ConnectDB()
	tx := db.MustBegin()
	_, err := tx.Exec(query, 4)
	tx.Commit()
	return err
}
