package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/middleware"
	db "task-process-service/internal/postgres"
	"task-process-service/internal/repository"
	"task-process-service/internal/service"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer(r *chi.Mux) {
	jwtService := service.NewJWTAuthService()
	authMiddleware := middleware.AuthMiddleware(*jwtService)
	authService := service.NewAuthService(repository.NewAuthRepository(db.Connection))
	taskService := service.NewTaskService(repository.NewTaskRepository(db.Connection))
	userService := service.NewUserService(repository.NewUserRepository(db.Connection))
	distributeTasksHandler := handler.NewDistributeTasksHandler(taskService, userService)

	r.Handle("/metrics", promhttp.Handler())

	r.HandleFunc("/distributeTasks", authMiddleware(middleware.RateLimit(distributeTasksHandler.DistributeTasks)))
	setupTaskRoutes(r, taskService, authMiddleware, middleware.RateLimit)
	setupAuthRoutes(r, authService)
	setupUserRoutes(r, userService, authMiddleware, middleware.RateLimit)

	http.ListenAndServe(":45009", r)
}
