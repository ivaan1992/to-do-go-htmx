package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"to-do-list/db"
)

type Task struct {
	Id   int
	Task string
	Done bool
}

//GET METHOD

func FetchTasks(w http.ResponseWriter, r *http.Request) {
	todos, _ := GetAllTasks(db.DB)

	tmpl.ExecuteTemplate(w, "taskList", todos)
}

//Utility functions

func GetAllTasks(dbPointer *sql.DB) ([]Task, error) {
	query := "SELECT * FROM tasks"

	rows, err := dbPointer.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var tasks []Task

	for rows.Next() {
		var todo Task
		rowErr := rows.Scan(&todo.Id, &todo.Task, &todo.Done)

		if rowErr != nil {
			return nil, rowErr
		}

		tasks = append(tasks, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskById(dbPointer *sql.DB, id int) (*Task, error) {
	query := "SELECT * FROM tasks WHERE id = ?"

	var task Task
	row := dbPointer.QueryRow(query, id)

	err := row.Scan(&task.Id, &task.Task, &task.Done)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No task was found with id %d", id)
		}

		return nil, err
	}

	return &task, nil
}
