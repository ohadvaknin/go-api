package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ohadvaknin/go-api/handlers"
)

func main() {
	router := mux.NewRouter()

	//Define routes
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}	