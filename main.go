package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/tasks", getTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", getTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", createTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", deleteTaskHandler).Methods("DELETE")

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
