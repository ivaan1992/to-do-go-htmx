package handlers

import (
	"net/http"
	"to-do-list/db"
	"to-do-list/models"
)

func FetchTasks(w http.ResponseWriter, r *http.Request) {
	database := db.DB
	defer database.Close()

	todos, _ := models.GetAllTasks(database)

	tmpl.ExecuteTemplate(w, "taskList", todos)
}
