package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ohadvaknin/go-api/handlers"
)
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Log request method, URI, and timestamp
        log.Printf("%s %s %s", r.Method, r.RequestURI, time.Now().Format(time.RFC1123))

        // Read and log the request body
        var bodyBytes []byte
        if r.Body != nil {
            bodyBytes, _ = io.ReadAll(r.Body)
        }

        r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore the request body
        log.Printf("Request Body: %s", string(bodyBytes))

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}

func main() {
	router := mux.NewRouter()

	//Define routes
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	loggedRouter := LoggingMiddleware(router)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}	