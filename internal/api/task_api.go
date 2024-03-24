package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/repository"
)

func setupTaskRoutes(repo repository.TaskRepository) {
	taskHandler := handler.NewTaskHandler(repo)
	http.HandleFunc("/tasks/add", taskHandler.Create)
	http.HandleFunc("/tasks/update", taskHandler.Update)
	http.HandleFunc("/tasks/delete", taskHandler.Delete)
	http.HandleFunc("/tasks/getById", taskHandler.GetById)
	http.HandleFunc("/tasks/getAll", taskHandler.GetAll)
	// Add other task routes here
}
