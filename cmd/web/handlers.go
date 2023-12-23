package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id <= 0 {
		http.Error(w, "id is not valid", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "This is the snippet that you are trying to view %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("This is the snippet create method"))
}

func snippetRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the root of the snippet"))
}
