package main

import (
	"net/http"
	"to-do-list/db"
	"to-do-list/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()
	defer db.DB.Close()

	gRouter := mux.NewRouter()

	gRouter.HandleFunc("/", handlers.HomeHandler)

	//Style Route
	gRouter.PathPrefix("/style/").Handler(http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))

	//Assets Route
	gRouter.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//Dark-Light Mode
	gRouter.HandleFunc("/toggle-theme", handlers.ToggleThemeHandler)

	//Tasks Handlers
	gRouter.HandleFunc("/tasks", handlers.FetchTasks).Methods("GET")

	//Run :3000 server
	http.ListenAndServe(":3000", gRouter)
}
