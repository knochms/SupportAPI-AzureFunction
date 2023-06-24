package handlers

import (
	"azure-function-support-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HandleTodosResponsibility(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleTodosResponsibility called.\n")
	url := r.URL.Path
	fmt.Fprint(w, url)

	url_split := strings.Split(url, "/")
	fmt.Fprint(w, url_split)
	responsibility := url_split[4]
	fmt.Fprint(w, responsibility)
	fmt.Fprint(w, url_split[3])
	fmt.Fprint(w, models.Todos)

	allTodosFromResponsibility := []models.Todo{}
	for _, t := range models.Todos {
		if t.Responsibility == responsibility {
			allTodosFromResponsibility = append(allTodosFromResponsibility, t)
		}
	}
	json.NewEncoder(w).Encode(allTodosFromResponsibility)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
