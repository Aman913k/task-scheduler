package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() error {
	var err error

	dsn := "root:Aman123@tcp(127.0.0.1:3306)/taskdb?parseTime=true"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	// Create tasks table if not exists
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INT AUTO_INCREMENT PRIMARY KEY,
			description VARCHAR(255) NOT NULL,
			status VARCHAR(50) NOT NULL DEFAULT 'pending',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err = db.Exec(query)
	return err
}

// Create a new task
func createTask(description string) (int, error) {
	result, err := db.Exec(
		"INSERT INTO tasks (description, status, created_at) VALUES (?, ?, ?)",
		description, "pending", time.Now(),
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

// Get a single task by ID
func getTask(id int) (Task, error) {
	var task Task
	row := db.QueryRow("SELECT id, description, status, created_at FROM tasks WHERE id = ?", id)
	err := row.Scan(&task.ID, &task.Description, &task.Status, &task.CreatedAt)
	return task, err
}

// Get all tasks
func getAllTasks() ([]Task, error) {
	rows, err := db.Query("SELECT id, description, status, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Status, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Update task status
func updateTaskStatus(id int, status string) error {
	_, err := db.Exec("UPDATE tasks SET status = ? WHERE id = ?", status, id)
	return err

}
