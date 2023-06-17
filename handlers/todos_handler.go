package handlers

import (
	"azure-function-support-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var todos []models.Todo = models.Todos
var history_all map[uuid.UUID][]models.Todo

func HandleTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleTodos called.")
	switch {
	case r.Method == http.MethodGet:
		// Handle GET /api/todos
		// ...
		getTodos(w, r)
	case r.Method == http.MethodPost:
		// Handle POST /api/todos
		// ...
		createTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleTodoById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HandleTodoById called.")
	vars := mux.Vars(r)
	id := vars["todoId"]

	// Verwende die id für die entsprechende Verarbeitung

	fmt.Fprintf(w, "HandleTodoById function called with id: %s", id)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getTodos called.")
	json.NewEncoder(w).Encode(todos)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createTodo called.")
	todo, err := models.ParseRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.CreatedAt = time.Now()
	todo.ModifiedAt = time.Now()
	todo.Id = uuid.New()
	todo.Assigned = false
	todo.Status = "created"
	todo.Responsibility = strings.ToLower(todo.Responsibility)
	todos = append(todos, *todo)

	models.Update_History(*todo)

	json.NewEncoder(w).Encode(todo.Id) //Rückgabeparameter, not used in Python-Support
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
