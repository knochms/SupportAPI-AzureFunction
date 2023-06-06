package handlers

import (
	"fmt"
	"net/http"
)

func HandleHistoryById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleHistoryById called.")
	// Implementieren Sie hier die Logik für den "/api/history/{todoId}" Endpunkt
}
