package services

import (
	"fmt"
	"net/http"
	"strconv"
	"stream-api/data"
	"stream-api/models"

	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

// GetUser Gets user.
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := data.GetUserByID(id)
	if err != nil {
		fmt.Println(err)
		resp := models.DoesNotExist{
			Code:    404,
			Message: "User does not exist",
		}
		return c.JSON(http.StatusNotFound, resp)
	}
	fmt.Println(user)

	return c.JSON(http.StatusOK, user)
}

// GetUsers Gets user.
func GetUsers(c echo.Context) error {
	user := data.GetUsers()
	fmt.Println(user)

	return c.JSON(http.StatusOK, user)
}

// CreateUser Creates a user
func CreateUser(c echo.Context) error {

	// Set form data
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")
	email := c.FormValue("email")

	u := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	fmt.Print("validating user")
	// Validate User
	result, err := valid.ValidateStruct(u)
	if err != nil {
		fmt.Println(err)
		ve := models.ValidationError{
			ValidationErrors: valid.ErrorsByField(err),
		}
		return c.JSON(http.StatusBadRequest, ve)
	}
	fmt.Println(result)

	// Create record
	user := data.CreateUser(u)
	return c.JSON(http.StatusCreated, user)
}

// UpdateUser Updates a user
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := data.GetUserByID(id)
	if err != nil {
		fmt.Println(err)
		resp := models.DoesNotExist{
			Code:    404,
			Message: "User does not exist",
		}
		return c.JSON(http.StatusNotFound, resp)
	}

	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")
	email := c.FormValue("email")

	if firstName != "" {
		user.FirstName = firstName
	}

	if lastName != "" {
		user.LastName = lastName
	}

	if email != "" {
		user.Email = email
	}

	updatedUser := data.UpdateUserByID(id, *user)
	return c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser deletes a user
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := data.DeleteUserByID(id)
	if err != nil {
		fmt.Println(err)
		resp := models.DoesNotExist{
			Code:    404,
			Message: "User does not exist",
		}
		return c.JSON(http.StatusNotFound, resp)
	}

	// resp := models.DoesNotExist{
	// 	Code:    200,
	// 	Message: "User deleted.",
	// }
	return c.NoContent(http.StatusNoContent)
}
