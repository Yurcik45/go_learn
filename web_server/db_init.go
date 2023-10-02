package main

import (
  "fmt"
  "log"
  "database/sql"

  _ "github.com/mattn/go-sqlite3"
)

func Init_db() {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "server_test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the "users" table
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			age INTEGER
		);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database and table created successfully!")
}
