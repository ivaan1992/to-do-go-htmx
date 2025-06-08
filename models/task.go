package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Task struct {
	Id           string
	Text         string
	Status       bool
	PendingCount string
}

func AddNewTask(db *sql.DB, task Task) error {
	query := "INSERT INTO tasks (task) VALUES (?)"
	_, err := db.Exec(query, task.Text)
	return err
}

func UpdateCurrentTask(db *sql.DB, id string, newTask string) error {
	query := "UPDATE tasks SET task = ? WHERE id = ?"

	result, err := db.Exec(query, newTask, id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("No se actualizó ninguna tarea con ID %d", id)
	}

	return nil
}

func DeleteCurrentTask(db *sql.DB, id string) error {
	query := "DELETE FROM tasks WHERE id = ?"

	result, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		fmt.Errorf("No task was deleted with id:", id)
	}

	return nil
}

func GetTaskById(db *sql.DB, id string) (*Task, error) {
	query := "SELECT * FROM tasks WHERE id = ?"

	var task Task
	row := db.QueryRow(query, id)

	err := row.Scan(&task.Id, &task.Text, &task.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No task was found with id %d", id)
		}

		return nil, err
	}

	return &task, nil
}

func CurrentToggleTask(db *sql.DB, id string, status bool) error {
	query := "UPDATE tasks SET done = ? WHERE id = ?"
	_, err := db.Exec(query, status, id)
	return err
}

func CountPendingTasks(db *sql.DB) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM tasks WHERE done = FALSE"
	err := db.QueryRow(query).Scan(&count)
	return count, err
}
