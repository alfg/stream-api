package controllers

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strconv"
	"stream-api/data"
	"stream-api/models"

	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

// GetStream Gets stream.
func GetStream(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	stream, err := data.GetStreamByID(id)
	if err != nil {
		fmt.Println(err)
		resp := models.DoesNotExist{
			Code:    404,
			Message: "Stream does not exist",
		}
		return c.JSON(http.StatusNotFound, resp)
	}
	fmt.Println(stream)

	return c.JSON(http.StatusOK, stream)
}

// GetStreams Gets stream.
func GetStreams(c echo.Context) error {
	stream := data.GetStreams()
	fmt.Println(stream)

	return c.JSON(http.StatusOK, stream)
}

// CreateStream Creates a stream
func CreateStream(c echo.Context) error {

	streamName := c.FormValue("stream_name")

	// Check if stream exists.
	if data.StreamExistsByName(streamName) {
		resp := models.AlreadyExists{
			Code:    409,
			Message: "Stream name is taken.",
		}
		return c.JSON(http.StatusConflict, resp)
	}

	// Set form data
	title := c.FormValue("title")
	_type := c.FormValue("type")
	description := c.FormValue("description")
	private, _ := strconv.ParseBool(c.FormValue("private"))

	streamKey := generateKey(10)

	s := models.Stream{
		Title:       title,
		Type:        _type,
		Description: description,
		Private:     private,
		StreamName:  streamName,
		StreamKey:   streamKey,
	}

	fmt.Print("validating stream")
	// Validate Stream
	result, err := valid.ValidateStruct(s)
	if err != nil {
		fmt.Println(err)
		ve := models.ValidationError{
			ValidationErrors: valid.ErrorsByField(err),
		}
		return c.JSON(http.StatusBadRequest, ve)
	}
	fmt.Println(result)

	// Create record
	stream := data.CreateStream(s)
	return c.JSON(http.StatusCreated, stream)
}

// UpdateStream Updates a stream
func UpdateStream(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	stream, err := data.GetStreamByID(id)
	if err != nil {
		fmt.Println(err)
		resp := models.DoesNotExist{
			Code:    404,
			Message: "Stream does not exist",
		}
		return c.JSON(http.StatusNotFound, resp)
	}

	title := c.FormValue("title")
	_type := c.FormValue("type")
	description := c.FormValue("email")
	private, _ := strconv.ParseBool(c.FormValue("private"))

	if title != "" {
		stream.Title = title
	}

	if _type != "" {
		stream.Type = _type
	}

	if description != "" {
		stream.Description = description
	}

	if private {
		stream.Private = private
	}

	updatedStream := data.UpdateStreamByID(id, *stream)
	return c.JSON(http.StatusOK, updatedStream)
}

// DeleteStream deletes a stream
func DeleteStream(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := data.DeleteStreamByID(id)
	if err != nil {
		fmt.Println(err)
		resp := models.DoesNotExist{
			Code:    404,
			Message: "Stream does not exist",
		}
		return c.JSON(http.StatusNotFound, resp)
	}

	// resp := models.StreamDoesNotExist{
	// 	Code:    200,
	// 	Message: "Stream deleted.",
	// }
	return c.NoContent(http.StatusNoContent)
}

// AuthenticateStream Authenticates stream by checking StreamKey.
func AuthenticateStream(c echo.Context) error {
	streamName := c.QueryParam("name")
	streamKey := c.QueryParam("key")
	if data.ValidateStreamKey(streamName, streamKey) {
		return c.String(http.StatusOK, "OK")
	}
	return c.String(http.StatusForbidden, "Forbidden")
}

func generateKey(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%x", b)
	return string(s)
}
