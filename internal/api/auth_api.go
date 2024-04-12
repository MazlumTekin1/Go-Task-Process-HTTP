package api

import (
	"task-process-service/internal/handler"
	"task-process-service/internal/service"

	"github.com/go-chi/chi"
)

func setupAuthRoutes(r *chi.Mux, auth service.AuthService) {
	authHandler := handler.NewAuthHandler(auth)
	r.HandleFunc("/login", authHandler.Login)
}
