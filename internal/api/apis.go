package api

import (
	"net/http"
	"task-process-service/docs"
	"task-process-service/internal/handler"
	"task-process-service/internal/middleware"
	db "task-process-service/internal/postgres"
	"task-process-service/internal/repository"
	"task-process-service/internal/service"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @title Go Task Process Service
// @version 1.0
// @description  Swagger documentation for the Go Task Process Service
// @host localhost:45009
// @BasePath /
func StartServer() {
	jwtService := service.NewJWTAuthService()
	authMiddleware := middleware.AuthMiddleware(*jwtService)
	authService := service.NewAuthService(repository.NewAuthRepository(db.Connection))
	taskService := service.NewTaskService(repository.NewTaskRepository(db.Connection))
	userService := service.NewUserService(repository.NewUserRepository(db.Connection))
	distributeTasksHandler := handler.NewDistributeTasksHandler(taskService, userService)

	r := chi.NewRouter()
	docs.SwaggerInfo.Title = "Go Task Process Service API Documentation"
	docs.SwaggerInfo.Description = "This is the API documentation for the Go Task Process Service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:45009"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:45009/swagger/doc.json"),
	))
	r.Handle("/metrics", promhttp.Handler())

	r.HandleFunc("/distributeTasks", authMiddleware(middleware.RateLimit(distributeTasksHandler.DistributeTasks)))
	setupTaskRoutes(r, taskService, authMiddleware, middleware.RateLimit)
	setupAuthRoutes(r, authService)
	setupUserRoutes(r, userService, authMiddleware, middleware.RateLimit)

	http.ListenAndServe(":45009", r)
}
