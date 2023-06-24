package main

import (
	"azure-function-support-api/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	router := mux.NewRouter()
	// Azure Functions hier
	router.HandleFunc("/api/todos/{todoId}", handlers.HandleTodoById)
	router.HandleFunc("/api/todos", handlers.HandleTodos)
	router.HandleFunc("/api/todo/responsibility/{responsibility}", handlers.HandleTodosResponsibility)
	router.HandleFunc("/api/todo/responsibility/{responsibility}/status/{status}", handlers.HandleTodosResponsibilityWithStatus)
	router.HandleFunc("/api/history/{todoId}", handlers.HandleHistoryById)

	//log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
