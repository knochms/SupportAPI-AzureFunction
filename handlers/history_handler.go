package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func HandleHistoryById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleHistoryById called.")
	urlPathSegments := strings.Split(r.URL.Path, "/")
	todoIDString := urlPathSegments[2]
	todoID, _ := uuid.Parse(todoIDString)
	for IdOfTodo, historieOfTodo := range history_all {
		if IdOfTodo == todoID {
			json.NewEncoder(w).Encode(historieOfTodo)
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
