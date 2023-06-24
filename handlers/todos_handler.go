package handlers

import (
	"azure-function-support-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
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
	switch {
	case r.Method == http.MethodGet:
		// Handle GET /api/todos/{todoID}
		// ...
		getTodo(w, r)
	case r.Method == http.MethodPut:
		// Handle PUT /api/todos/{todoID}
		// ...
		updateTodo(w, r)
	case r.Method == http.MethodDelete:
		// Handle DELETE /api/todos/{todoID}
		// ...
		deleteTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	/* vars := mux.Vars(r)
	id := vars["todoId"]
	// Verwende die id für die entsprechende Verarbeitung
	fmt.Fprintf(w, "HandleTodoById function called with id: %s", id) */
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

func getTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getTodo called.")
	urlPathSegments := strings.Split(r.URL.Path, "/")

	todoIDString := urlPathSegments[3]
	todoID, err := uuid.Parse(todoIDString)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	var foundToDo models.Todo

	for _, t := range todos {
		if t.Id == todoID {
			foundToDo = t
		}
	}
	fmt.Println(foundToDo)
	json.NewEncoder(w).Encode(foundToDo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "updateTodo called.")
	urlPathSegments := strings.Split(r.URL.Path, "/")
	todoIDString := urlPathSegments[3]
	todoID, err := uuid.Parse(todoIDString)
	fmt.Println(todoID)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	todo, err := models.ParseRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, t := range todos {
		if t.Id == todoID {
			todos[i].Title = todo.Title
			todos[i].Description = todo.Description
			todos[i].Status = todo.Status
			todos[i].Priority = todo.Priority
			todos[i].ModifiedAt = time.Now()
			todos[i].Responsibility = strings.ToLower(todo.Responsibility)

			if todos[i].Status == "done" {
				todos[i].CompletedAt = time.Now()
			}

			update_History(todos[i])

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "deleteTodo called.")
	urlPathSegments := strings.Split(r.URL.Path, "/")
	todoIDString := urlPathSegments[3]
	todoID, err := uuid.Parse(todoIDString)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, t := range todos {
		if t.Id == todoID {
			todos[i].Status = "deleted"
			todos[i].ModifiedAt = time.Now()
			update_History(todos[i])
			break
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func update_History(history_entry models.Todo) {
	if history_all == nil {
		history_all = make(map[uuid.UUID][]models.Todo)
	}
	_, exists := history_all[history_entry.Id]
	if !exists {
		history_all[history_entry.Id] = make([]models.Todo, 0)
	}
	history_all[history_entry.Id] = append(history_all[history_entry.Id], history_entry)
	fmt.Println(history_all)
}
