package handlers

import (
	"fmt"
	"net/http"
)

func HandleTodosResponsibility(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleTodosResponsibility called.")
	/* responsibility := r.URL.Query().Get("responsibility")
	// Case: /api/todo/responsibility/{responsibility}
	if responsibility != "" {
		fmt.Println("HandleTodosResponsibility function called")
	} */
	/* 	switch {
	   	case r.Method == http.MethodGet && responsibility != "":
	   		// Handle GET /api/todo/responsibility/{responsibility}
	   		// ...
	   	default:
	   		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	   	} */
}
