package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/service"
)

func setupAuthRoutes(auth service.AuthService) {
	authHandler := handler.NewAuthHandler(auth)
	http.HandleFunc("/login", authHandler.Login)
}
