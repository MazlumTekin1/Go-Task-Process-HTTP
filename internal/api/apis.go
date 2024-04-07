package api

import (
	"net/http"
	"task-process-service/internal/handler"
	"task-process-service/internal/middleware"
	db "task-process-service/internal/postgres"
	"task-process-service/internal/repository"
	"task-process-service/internal/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

func StartServer() {

	jwtService := service.NewJWTAuthService()

	authMiddleware := middleware.AuthMiddleware(*jwtService)
	http.Handle("/metrics", promhttp.Handler())
	authService := service.NewAuthService(repository.NewAuthRepository(db.Connection))
	taskService := service.NewTaskService(repository.NewTaskRepository(db.Connection))
	userService := service.NewUserService(repository.NewUserRepository(db.Connection))
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	distributeTasksHandler := handler.NewDistributeTasksHandler(taskService, userService)

	http.HandleFunc("/distributeTasks", authMiddleware(middleware.RateLimit(distributeTasksHandler.DistributeTasks)))
	setupTaskRoutes(taskService, authMiddleware, middleware.RateLimit)
	setupAuthRoutes(authService)
	setupUserRoutes(userService, authMiddleware, middleware.RateLimit)

	http.ListenAndServe(":45009", nil)
}
