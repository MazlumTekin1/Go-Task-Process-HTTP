package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/service"
)

func setupTaskRoutes(s service.TaskService, authMiddleware func(http.HandlerFunc) http.HandlerFunc, rateLimit func(http.HandlerFunc) http.HandlerFunc) {
	taskHandler := handler.NewTaskHandler(s)
	http.HandleFunc("/tasks/add", authMiddleware(taskHandler.Create))
	http.HandleFunc("/tasks/update", authMiddleware(taskHandler.Update))
	http.HandleFunc("/tasks/delete", authMiddleware(taskHandler.Delete))
	http.HandleFunc("/tasks/getById", authMiddleware(taskHandler.GetById))
	http.HandleFunc("/tasks/getAll", authMiddleware(rateLimit(taskHandler.GetAll)))
}
