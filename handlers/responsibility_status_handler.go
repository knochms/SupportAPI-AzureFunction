package handlers

import (
	"fmt"
	"net/http"
)

func HandleTodosResponsibilityWithStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleTodosResponsibilityWithStatus called.")
	/* responsibility := r.URL.Query().Get("responsibility")
	status := r.URL.Query().Get("status")
	if responsibility != "" && status != "" {
		fmt.Println("HandleTodosResponsibilityWithStatus function called")
	} */
	/* switch {
	case r.Method == http.MethodGet && responsibility != "" && status != "":
		// Handle GET /api/todo/responsibility/{responsibility}/status/{status}
		// ...
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	} */
}
