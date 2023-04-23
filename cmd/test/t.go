package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func main() {
	http.HandleFunc("/todos", handleTodos)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		createTodo(w, r)
	case http.MethodPut:
		updateTodo(w, r)
	case http.MethodDelete:
		deleteTodo(w, r)
	default:
		http.Error(w, fmt.Sprintf("Invalid HTTP method: %s", r.Method), http.StatusBadRequest)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.ID = len(todos) + 1
	todos = append(todos, todo)

	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	if idString == "" {
		http.Error(w, "ID query parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID: %s", idString), http.StatusBadRequest)
		return
	}

	var todo Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	found := false
	for i, t := range todos {
		if t.ID == id {
			todos[i] = todo
			found = true
			break
		}
	}

	if !found {
		http.Error(w, fmt.Sprintf("Todo with ID %d not found", id), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	if idString == "" {
		http.Error(w, "ID query parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID: %s", idString), http.StatusBadRequest)
		return
	}

	found := false
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, fmt.Sprintf("Todo with ID %d not found", id), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
