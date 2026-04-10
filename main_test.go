package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/tasks", getTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", getTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", createTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", deleteTaskHandler).Methods("DELETE")
	return r
}

func TestHealthHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	setupRouter().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Code)
	}
}

func TestCreateAndGetTask(t *testing.T) {
	tasks = []Task{}
	nextID = 1

	body, _ := json.Marshal(Task{Title: "Write Dockerfile"})
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	setupRouter().ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected 201 got %d", rr.Code)
	}

	req2, _ := http.NewRequest("GET", "/tasks", nil)
	rr2 := httptest.NewRecorder()
	setupRouter().ServeHTTP(rr2, req2)

	var result []Task
	if err := json.NewDecoder(rr2.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("expected 1 task got %d", len(result))
	}
}
func TestDeleteTask(t *testing.T) {
	tasks = []Task{{ID: 1, Title: "Test task", Done: false}}
	nextID = 2

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	rr := httptest.NewRecorder()
	setupRouter().ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("expected 204 got %d", rr.Code)
	}
}
