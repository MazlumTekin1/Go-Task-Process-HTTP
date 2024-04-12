package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/service"

	"github.com/go-chi/chi"
)

func setupUserRoutes(r *chi.Mux, s service.UserService, authMiddleware func(http.HandlerFunc) http.HandlerFunc, rateLimit func(http.HandlerFunc) http.HandlerFunc) {
	userHandler := handler.NewUserHandler(s)
	r.HandleFunc("/users/add", authMiddleware(rateLimit(userHandler.Create)))
	r.HandleFunc("/users/update", authMiddleware(rateLimit(userHandler.Update)))
	r.HandleFunc("/users/delete", authMiddleware(rateLimit(userHandler.Delete)))
	r.HandleFunc("/users/getById", authMiddleware(rateLimit(userHandler.GetById)))
	r.HandleFunc("/users/getAll", authMiddleware(rateLimit(userHandler.GetAll)))
}
