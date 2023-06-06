package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createTodo called.")
}
