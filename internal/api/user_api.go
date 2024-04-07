package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/service"
)

func setupUserRoutes(s service.UserService, authMiddleware func(http.HandlerFunc) http.HandlerFunc) {
	userHandler := handler.NewUserHandler(s)
	http.HandleFunc("/users/add", authMiddleware(userHandler.Create))
	http.HandleFunc("/users/update", authMiddleware(userHandler.Update))
	http.HandleFunc("/users/delete", authMiddleware(userHandler.Delete))
	http.HandleFunc("/users/getById", authMiddleware(userHandler.GetById))
	http.HandleFunc("/users/getAll", authMiddleware(userHandler.GetAll))
}
