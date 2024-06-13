package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/ohadvaknin/go-api/models"
)

// In-memory storage for tasks
var tasks = make(map[string]models.Task)

// AddTask adds a new task to the storage
func AddTask(task models.Task) models.Task {
    task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	tasks[task.ID] = task
	return task
}

// GetAllTasks returns all tasks from the storage
func GetAllTasks() []models.Task {
    // Implementation here
	var taskList []models.Task
	for _,v := range tasks {
		taskList = append(taskList, v)
	}
	return taskList
}

// GetTaskByID returns a task by its ID from the storage
func GetTaskByID(id string) (models.Task, bool) {
    task, exists := tasks[id]
	return task, exists
}

// UpdateTask updates an existing task in the storage
func UpdateTask(id string, updatedTask models.Task) (models.Task, bool) {
    task, exists := tasks[id]
	if !exists {
		return models.Task{}, false
	} 
	updatedTask.ID = task.ID
	updatedTask.CreatedAt = task.CreatedAt
	updatedTask.UpdatedAt = time.Now()
	tasks[updatedTask.ID] = updatedTask
	return updatedTask, true
}

// DeleteTask deletes a task by its ID from the storage
func DeleteTask(id string) bool {
	_, exists := tasks[id]
	if !exists {
		return false
	}
    delete(tasks, id)
	return true
}
