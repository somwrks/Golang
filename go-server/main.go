package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type FormData struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

var formSubmissions []FormData
var mu sync.Mutex

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusBadRequest)
		log.Printf("Error parsing form: %v\n", err)
		return
	}

	formData := FormData{
		Name:    r.FormValue("name"),
		Address: r.FormValue("address"),
		Email:   r.FormValue("email"),
		Phone:   r.FormValue("phone"),
	}

	mu.Lock()
	formSubmissions = append(formSubmissions, formData)
	mu.Unlock()

	log.Printf("Received submission: %v\n", formData)
	fmt.Fprintf(w, "Form submission successful!\n")
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/summaryData" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	json.NewEncoder(w).Encode(formSubmissions)
	mu.Unlock()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/summaryData", summaryHandler)
	http.HandleFunc("/hello", helloHandler)

	// Setup logging to a file
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("Starting server at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
