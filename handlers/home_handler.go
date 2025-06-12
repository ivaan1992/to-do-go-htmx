package handlers

import (
	"net/http"
	"text/template"
	"to-do-list/db"
	"to-do-list/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

type PageData struct {
	Theme ThemeData
	Tasks []models.Task
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetAllTasks(db.DB)
	if err != nil {
		http.Error(w, "Error fetching tasks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Theme: GetCurrentTheme(),
		Tasks: tasks,
	}

	err = tmpl.ExecuteTemplate(w, "Index.html", data)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
