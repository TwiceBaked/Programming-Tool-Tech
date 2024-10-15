package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
	"embed"
	"io/fs"
)

func handlerTEXT(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World!")
}

func loadFile(fileName string)(string, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handlerHTML(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Hello World GWS</h1>"))
}

func help(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var html, err = loadFile("help.html")
	if err != nil {
        http.Error(w, "Could not load help page", http.StatusInternalServerError)
        return
    }
	fmt.Fprint(w, html)
}

type Response struct {
	Message string "json:message"
}
func handlerJSON(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    response := Response{Message: "Hello, this is JSON!"}
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

//go:embed syllabus.json
var syllabus embed.FS

func handlerSYLLABUS(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		data, err := fs.ReadFile(syllabus, "syllabus.json")
		if err != nil {
			http.Error(w, "Failed to read json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func deleteSyllabus(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("deleted - stubbed"))
}

func create(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
    w.Write([]byte("create - stubbed"))
}

func update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
    w.Write([]byte("update - stubbed"))
}

func main() {
	http.HandleFunc("/", handlerTEXT)
	http.HandleFunc("/html", handlerHTML)
	http.HandleFunc("/json", handlerJSON)
	http.HandleFunc("/syllabus", handlerSYLLABUS)
	http.HandleFunc("/syllabus/delete", deleteSyllabus)
	http.HandleFunc("/syllabus/create", create)
	http.HandleFunc("/syllabus/update", update)
	http.HandleFunc("/help", help)
	http.ListenAndServe(":9000", nil)
}