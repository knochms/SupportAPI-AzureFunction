package handlers

import (
	"azure-function-support-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func HandleHistoryById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleHistoryById called.\n")
	urlPathSegments := strings.Split(r.URL.Path, "/")
	todoIDString := urlPathSegments[3]
	todoID, _ := uuid.Parse(todoIDString)
	for IdOfTodo, historieOfTodo := range models.History_all {
		if IdOfTodo == todoID {
			json.NewEncoder(w).Encode(historieOfTodo)
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
