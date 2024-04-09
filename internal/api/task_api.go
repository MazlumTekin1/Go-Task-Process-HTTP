package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/service"
)

func setupTaskRoutes(s service.TaskService, authMiddleware func(http.HandlerFunc) http.HandlerFunc, rateLimit func(http.HandlerFunc) http.HandlerFunc) {
	taskHandler := handler.NewTaskHandler(s)
	http.HandleFunc("/tasks/add", authMiddleware(rateLimit(taskHandler.Create)))
	http.HandleFunc("/tasks/update", authMiddleware(rateLimit(taskHandler.Update)))
	http.HandleFunc("/tasks/delete", authMiddleware(rateLimit(taskHandler.Delete)))
	http.HandleFunc("/tasks/getById", authMiddleware(rateLimit(taskHandler.GetById)))
	http.HandleFunc("/tasks/getAll", authMiddleware(rateLimit(taskHandler.GetAll)))
}
