package db

// The underscore indicates that the package is not directly but indirecly used
import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB = nil

func InitDB() {
	var err error
	// A sqlite db uses a local file
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		// If there is no db then it doesn't make sense to continue.
		// Then we crash the app
		// Gin might avoid the app crashing and try to recover
		panic("Could not connect to database.")
	}

	// How many connections we are going to work with
	DB.SetMaxOpenConns(10)
	// How many connections open if non are being used
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER, 
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)
	fmt.Println(err)
	if err != nil {
		panic("Could not create events table.")
	}

}
