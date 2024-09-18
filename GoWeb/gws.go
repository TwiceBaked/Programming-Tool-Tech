package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
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
	var html, _ = loadFile("GoWeb/index.html")
	fmt.Fprintf(w, html)
	
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



func main() {
	http.HandleFunc("/", handlerTEXT)
	http.HandleFunc("/html", handlerHTML)
	http.HandleFunc("/json", handlerJSON)
	http.ListenAndServe(":9000", nil)
	
}