package handlers

import (
	"net/http"
	"to-do-list/db"
	"to-do-list/models"

	"github.com/gorilla/mux"
)

// SHOW Tasks
func FetchTasks(w http.ResponseWriter, r *http.Request) {
	database := db.DB

	todos, err := models.GetAllTasks(database)

	if err != nil {
		http.Error(w, "Error Fetching Tasks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := models.Tasks{
		Tasks: todos,
	}

	err = tmpl.ExecuteTemplate(w, "TaskList", data)

	if err != nil {
		http.Error(w, "Error Rendering TaskList: "+err.Error(), http.StatusInternalServerError)
	}
}

//ADD Tasks

func CreateTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	text := r.FormValue("TaskInput")

	if text == "" {
		http.Error(w, "Task text is required", http.StatusBadRequest)
		return
	}

	task := models.Task{
		Text:   text,
		Status: false,
	}

	err = task.Save(db.DB)
	if err != nil {
		http.Error(w, "Error inserting task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tasks, _ := models.GetAllTasks(db.DB)
	tmpl.ExecuteTemplate(w, "TaskList", models.Tasks{Tasks: tasks})
}

// UPDATE Tasks

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	taskText := r.FormValue("TaskInput")

	vars := mux.Vars(r)
	taskId := vars["id"]

	task := models.Task{
		Id:   taskId,
		Text: taskText,
	}

	err = task.Update(db.DB)

	if err != nil {
		http.Error(w, "Error actualizando tarea: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tasks, _ := models.GetAllTasks(db.DB)
	tmpl.ExecuteTemplate(w, "TaskList", models.Tasks{Tasks: tasks})
}

func ShowUpdateForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	task := models.Task{
		Id: taskId,
	}

	err := task.Read(db.DB)
	if err != nil {
		http.Error(w, "Tarea no encontrada: "+err.Error(), http.StatusNotFound)
		return
	}

	err = tmpl.ExecuteTemplate(w, "UpdateTaskForm", task)
	if err != nil {
		http.Error(w, "Error al renderizar el formulario: "+err.Error(), http.StatusInternalServerError)
	}
}

// DELETE Task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	task := models.Task{
		Id: taskId,
	}

	err := task.Delete(db.DB)
	if err != nil {
		http.Error(w, "Error al eliminar la tarea: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tasks, _ := models.GetAllTasks(db.DB)
	tmpl.ExecuteTemplate(w, "TaskList", models.Tasks{Tasks: tasks})
}

// Toggle
func ToggleTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	status := r.FormValue("status") == "on"

	task := models.Task{
		Id:     taskId,
		Status: status,
	}

	err := task.Check(db.DB)
	if err != nil {
		http.Error(w, "Error al cambiar estado: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tasks, err := models.GetAllTasks(db.DB)
	if err != nil {
		http.Error(w, "Error al obtener tareas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "TaskList", models.Tasks{Tasks: tasks})
}
