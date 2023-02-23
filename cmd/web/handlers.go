package main

import (
	"net/http"
)

// Create our handler functions
func (app *application) Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my webpage"))
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//w.Write([]byte("This is my home"))
}

func (app *application) About(w http.ResponseWriter, r *http.Request) {
	//day := time.Now().Weekday()
	//w.Write([]byte(fmt.Sprintf("Have a good %s.", day)))
}

func (app *application) MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//w.Write([]byte("Method not allowed"))
		w.Header().Set("Allow", "POST")
		//w.WriteHeader(405)
		http.Error(w, "Methold not allowed", http.StatusMethodNotAllowed)
		return
	}
	//w.Write([]byte("Message created..."))
}
