package handlers

import (
	"azure-function-support-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HandleTodosResponsibilityWithStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleTodosResponsibilityWithStatus called.")
	url := r.URL.Path

	url_split := strings.Split(url, "/")
	responsibility := url_split[3]
	status := url_split[5]

	allTodosFromResponsibilitywithStatus := []models.Todo{}
	for _, t := range todos {
		if t.Responsibility == responsibility && t.Status == status {
			allTodosFromResponsibilitywithStatus = append(allTodosFromResponsibilitywithStatus, t)
		}
	}
	json.NewEncoder(w).Encode(allTodosFromResponsibilitywithStatus)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
