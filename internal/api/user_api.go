package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/service"
)

func setupUserRoutes(s service.UserService, authMiddleware func(http.HandlerFunc) http.HandlerFunc, rateLimit func(http.HandlerFunc) http.HandlerFunc) {
	userHandler := handler.NewUserHandler(s)
	http.HandleFunc("/users/add", authMiddleware(rateLimit(userHandler.Create)))
	http.HandleFunc("/users/update", authMiddleware(rateLimit(userHandler.Update)))
	http.HandleFunc("/users/delete", authMiddleware(rateLimit(userHandler.Delete)))
	http.HandleFunc("/users/getById", authMiddleware(rateLimit(userHandler.GetById)))
	http.HandleFunc("/users/getAll", authMiddleware(rateLimit(userHandler.GetAll)))
}
