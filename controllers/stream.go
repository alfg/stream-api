package controllers

import (
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

	// Set form data
	streamName := c.FormValue("stream_name")
	_type := c.FormValue("type")
	description := c.FormValue("email")
	url := c.FormValue("url")
	private, _ := strconv.ParseBool(c.FormValue("private"))

	u := models.Stream{
		StreamName:  streamName,
		Type:        _type,
		Description: description,
		URL:         url,
		Private:     private,
	}

	fmt.Print("validating stream")
	// Validate Stream
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
	stream := data.CreateStream(u)
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

	streamName := c.FormValue("stream_name")
	_type := c.FormValue("type")
	description := c.FormValue("email")
	url := c.FormValue("url")
	private, _ := strconv.ParseBool(c.FormValue("private"))

	if streamName != "" {
		stream.StreamName = streamName
	}

	if _type != "" {
		stream.Type = _type
	}

	if description != "" {
		stream.Description = description
	}

	if url != "" {
		stream.URL = url
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
