package handlers

import (
	"net/http"
	"text/template"
	"to-do-list/db"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

type PageData struct {
	Theme ThemeData
	Tasks []Task
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := GetAllTasks(db.DB)
	if err != nil {
		http.Error(w, "Error fetching tasks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Theme: GetCurrentTheme(),
		Tasks: tasks,
	}

	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
