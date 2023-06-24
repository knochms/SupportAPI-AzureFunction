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

//var todos []models.Todo = models.Todos
//var history_all map[uuid.UUID][]models.Todo

func HandleTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleTodos called.\n")
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
	fmt.Fprintf(w, "HandleTodoById called.\n")
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
	fmt.Fprint(w, "getTodos called.\n")
	json.NewEncoder(w).Encode(models.Todos)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createTodo called.\n")
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
	models.Todos = append(models.Todos, *todo)

	models.Update_History(*todo)

	json.NewEncoder(w).Encode(todo.Id) //Rückgabeparameter, not used in Python-Support
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getTodo called.\n")
	urlPathSegments := strings.Split(r.URL.Path, "/")

	todoIDString := urlPathSegments[3]
	todoID, err := uuid.Parse(todoIDString)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	var foundToDo models.Todo

	for _, t := range models.Todos {
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
	fmt.Fprint(w, "updateTodo called.\n")
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
	for i, t := range models.Todos {
		if t.Id == todoID {
			models.Todos[i].Title = todo.Title
			models.Todos[i].Description = todo.Description
			models.Todos[i].Status = todo.Status
			models.Todos[i].Priority = todo.Priority
			models.Todos[i].ModifiedAt = time.Now()
			models.Todos[i].Responsibility = strings.ToLower(todo.Responsibility)

			if models.Todos[i].Status == "done" {
				models.Todos[i].CompletedAt = time.Now()
			}

			update_History(models.Todos[i])

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "deleteTodo called.\n")
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

	for i, t := range models.Todos {
		if t.Id == todoID {
			models.Todos[i].Status = "deleted"
			models.Todos[i].ModifiedAt = time.Now()
			update_History(models.Todos[i])
			break
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func update_History(history_entry models.Todo) {
	if models.History_all == nil {
		models.History_all = make(map[uuid.UUID][]models.Todo)
	}
	_, exists := models.History_all[history_entry.Id]
	if !exists {
		models.History_all[history_entry.Id] = make([]models.Todo, 0)
	}
	models.History_all[history_entry.Id] = append(models.History_all[history_entry.Id], history_entry)
	fmt.Println(models.History_all)
}
