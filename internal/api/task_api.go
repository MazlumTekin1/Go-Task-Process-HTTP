package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/service"

	"github.com/go-chi/chi"
)

func setupTaskRoutes(r *chi.Mux, s service.TaskService, authMiddleware func(http.HandlerFunc) http.HandlerFunc, rateLimit func(http.HandlerFunc) http.HandlerFunc) {
	taskHandler := handler.NewTaskHandler(s)
	r.HandleFunc("/tasks/add", authMiddleware(rateLimit(taskHandler.Create)))
	r.HandleFunc("/tasks/update", authMiddleware(rateLimit(taskHandler.Update)))
	r.HandleFunc("/tasks/delete", authMiddleware(rateLimit(taskHandler.Delete)))
	r.HandleFunc("/tasks/getById/{id}", authMiddleware(rateLimit(taskHandler.GetById)))
	r.HandleFunc("/tasks/getAll", authMiddleware(rateLimit(taskHandler.GetAll)))
	r.HandleFunc("/tasks/taskStatus", authMiddleware(rateLimit(taskHandler.GetAllTaskStatus)))
}
