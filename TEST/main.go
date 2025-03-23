package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create a database object which can be used to connect with the database.
	db, err := sql.Open("mysql", "root@tcp(0.0.0.0:3306)/field")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// Ensure the database connection is closed when the function ends.
	defer db.Close()

	// Check the database connection.
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Create the table with a properly formatted SQL query.
	createTableQuery := `
	CREATE TABLE user (
		id INT NOT NULL,
		name VARCHAR(20),
		PRIMARY KEY (id)
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	fmt.Println("Successfully Created Table")
}
