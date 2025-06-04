package models

import (
	"database/sql"
)

type Tasks struct {
	Tasks []Task
}

func GetAllTasks(db *sql.DB) ([]Task, error) {
	query := "SELECT * FROM tasks"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {

		var task Task
		rowErr := rows.Scan(&task.Id, &task.Text, &task.Status)

		if rowErr != nil {
			return nil, rowErr
		}

		tasks = append(tasks, task)

	}
	return tasks, nil
}
