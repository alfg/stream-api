package data

import (
	"fmt"
	"stream-api/models"
)

// Stream model

// GetStreamByID Gets stream by Id
func GetStreamByID(id int) (*models.Stream, error) {
	const query = `SELECT * FROM stream WHERE id = $1`

	db, _ := ConnectDB()
	stream := models.Stream{}
	err := db.Get(&stream, query, id)
	fmt.Println(&stream)
	if err != nil {
		fmt.Println(err)
		return &stream, err
	}

	return &stream, nil
}

// GetStreams Gets all streams.
func GetStreams() *[]models.Stream {
	const query = `SELECT * FROM stream ORDER BY id ASC`

	db, _ := ConnectDB()
	stream := []models.Stream{}
	db.Select(&stream, query)
	fmt.Println(&stream)

	return &stream
}

// CreateStream creates a stream.
func CreateStream(stream models.Stream) *models.Stream {
	const query = "INSERT INTO stream (stream_name, type, description, url, key, private) VALUES (:stream_name, :type, :description, :url, :key, :private)"

	db, _ := ConnectDB()
	tx := db.MustBegin()
	result, err := tx.NamedExec(query, &stream)
	if err != nil {
		fmt.Println("Error", err)
	}
	tx.Commit()

	fmt.Println("transaction done")

	lastID, _ := result.LastInsertId()
	stream.ID = lastID

	return &stream
}

// UpdateStreamByID Update stream by id
func UpdateStreamByID(id int, stream models.Stream) *models.Stream {
	const query = `UPDATE stream
		SET stream_name = :stream_name,
		type = :type,
		description = :description
		url = :url
		key = :key
		private = :private
		WHERE id = :id`

	db, _ := ConnectDB()
	tx := db.MustBegin()
	_, err := tx.NamedExec(query, &stream)
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()

	return &stream
}

// DeleteStreamByID Deletes stream by id
func DeleteStreamByID(id int) error {
	const query = "DELETE FROM stream WHERE id = :id"

	db, _ := ConnectDB()
	tx := db.MustBegin()
	_, err := tx.Exec(query, 4)
	tx.Commit()
	return err
}
