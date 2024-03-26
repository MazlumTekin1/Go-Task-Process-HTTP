package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/service"
)

func setupTaskRoutes(s service.TaskService) {
	taskHandler := handler.NewTaskHandler(s)
	http.HandleFunc("/tasks/add", taskHandler.Create)
	http.HandleFunc("/tasks/update", taskHandler.Update)
	http.HandleFunc("/tasks/delete", taskHandler.Delete)
	http.HandleFunc("/tasks/getById", taskHandler.GetById)
	http.HandleFunc("/tasks/getAll", taskHandler.GetAll)
}
