package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/service"
)

func setupUserRoutes(s service.UserService) {
	userHandler := handler.NewUserHandler(s)
	http.HandleFunc("/users/add", userHandler.Create)
	http.HandleFunc("/users/update", userHandler.Update)
	http.HandleFunc("/users/delete", userHandler.Delete)
	http.HandleFunc("/users/getById", userHandler.GetById)
	http.HandleFunc("/users/getAll", userHandler.GetAll)
}
