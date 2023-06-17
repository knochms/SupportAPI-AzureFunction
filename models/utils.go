package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var Todos []Todo
var History_all map[uuid.UUID][]Todo

func GetTodos() []Todo {
	return Todos
}

func SetTodos(todos_updated []Todo) {
	Todos = todos_updated
}

func ParseRequest(w http.ResponseWriter, r *http.Request) (*Todo, error) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &todo, nil
}

func Update_History(history_entry Todo) {
	if History_all == nil {
		History_all = make(map[uuid.UUID][]Todo)
	}
	_, exists := History_all[history_entry.Id]
	if !exists {
		History_all[history_entry.Id] = make([]Todo, 0)
	}
	History_all[history_entry.Id] = append(History_all[history_entry.Id], history_entry)
	fmt.Println(History_all)
}
