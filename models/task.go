package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Task struct {
	Id     string
	Text   string
	Status bool
}

func (t *Task) Save(db *sql.DB) error {
	query := "INSERT INTO tasks (task) VALUES (?)"
	_, err := db.Exec(query, t.Text)
	return err
}

func (t *Task) Update(db *sql.DB) error {
	query := "UPDATE tasks SET task = ? WHERE id = ?"

	result, err := db.Exec(query, t.Text, t.Id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("No task was updated with id: %s", t.Id)
	}

	return nil
}

func (t *Task) Delete(db *sql.DB) error {
	query := "DELETE FROM tasks WHERE id = ?"

	result, err := db.Exec(query, t.Id)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("error checking affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No task was deleted with id:", t.Id)
	}

	return nil
}

func (t *Task) Read(db *sql.DB) error {
	query := "SELECT * FROM tasks WHERE id = ?"
	row := db.QueryRow(query, t.Id)

	err := row.Scan(&t.Id, &t.Text, &t.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("No task was found with id %s", t.Id)
		}
		return err
	}

	return nil
}

func (t *Task) Check(db *sql.DB) error {
	query := "UPDATE tasks SET done = ? WHERE id = ?"
	_, err := db.Exec(query, t.Status, t.Id)
	return err
}
