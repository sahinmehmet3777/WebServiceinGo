// database.go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// getDB returns a database connection
func getDB() (*sql.DB, error) {
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", "user=username password=password dbname=mydb sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	return db, nil
}

// createTask creates a new task in the database
func createTask(title, description, status string) (int, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// Insert a new task into the "tasks" table
	var taskID int
	err = db.QueryRow("INSERT INTO tasks(title, description, status) VALUES($1, $2, $3) RETURNING id", title, description, status).Scan(&taskID)
	if err != nil {
		return 0, fmt.Errorf("error inserting task: %v", err)
	}

	return taskID, nil
}

// getTasks returns a list of tasks from the database
func getTasks() ([]Task, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query all tasks from the "tasks" table
	rows, err := db.Query("SELECT id, title, description, status FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("error querying tasks: %v", err)
	}
	defer rows.Close()

	// Iterate through the rows and create Task objects
	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return nil, fmt.Errorf("error scanning task: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
