package models

import (
	"time"

	"example.com/REST-API/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

// var events = []Event{}

func (e *Event) Save() error {
	// The question marks give a SQL injection safe way to make a query
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	// Prepare can help to get better performance than just using Exec()
	// It helpt keeping the query in memory to use it again
	// That is only true if we use it before we close it
	stmt, err := db.DB.Prepare(query)

	if err != nil {

		return err
	}

	defer stmt.Close()

	// Exec() is used when you modify stuff in the db
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	// This is the other way to use a query
	// Query() is typically used when you expect to get data
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	// Next returns a boolean indicating if there is a next one or not
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}
