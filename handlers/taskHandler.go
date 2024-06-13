package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ohadvaknin/go-api/models"
	"github.com/ohadvaknin/go-api/storage"
)

// CreateTask handles the creation of a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    err := json.NewDecoder(r.Body).Decode(&task)
    if (err != nil) {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    task = storage.AddTask(task)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

// GetTasks handles fetching all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
    taskList := storage.GetAllTasks()
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(taskList)
    if (err != nil) {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

}

// GetTask handles fetching a specific task by its ID
func GetTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    task, exists := storage.GetTaskByID(id)
    if !exists {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(task)
    if (err != nil) {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// UpdateTask handles updating an existing task by its ID
func UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var updatedTask models.Task
    errBody := json.NewDecoder(r.Body).Decode(&updatedTask)
    if (errBody != nil) {
        http.Error(w, errBody.Error(), http.StatusBadRequest)
        return
    }

    task, exists := storage.UpdateTask(id, updatedTask)
    if (!exists) {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(task)
    if (err != nil) {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

}

// DeleteTask handles deleting a task by its ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    isDeleted := storage.DeleteTask(id)
    if !isDeleted {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
