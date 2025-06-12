package handlers

import (
	"net/http"
	"sync"
	"text/template"
)

type ThemeData struct {
	BgImage string
	Icon    string
}

var (
	darkMode = false
	mu       sync.RWMutex
)

func GetCurrentTheme() ThemeData {
	mu.RLock()
	defer mu.RUnlock()

	data := ThemeData{}
	if darkMode {
		data.BgImage = "/assets/images/bg-desktop-dark.jpg"
		data.Icon = "/assets/icons/icon-sun.svg"
	} else {
		data.BgImage = "/assets/images/bg-desktop-light.jpg"
		data.Icon = "/assets/icons/icon-moon.svg"
	}
	return data
}

func ToggleThemeHandler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	darkMode = !darkMode
	mu.Unlock()

	data := GetCurrentTheme()

	tmpl := template.Must(template.ParseFiles("templates/Header.html"))
	err := tmpl.ExecuteTemplate(w, "Header", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
