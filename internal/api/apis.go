package api

import (
	"net/http"
	postgres "task-process-service/internal/infrastructure"
	"task-process-service/internal/repository"

	httpSwagger "github.com/swaggo/http-swagger"
)

func StartServer() {
	taskRepo := repository.NewTaskRepository(postgres.Connection)
	userRepo := repository.NewUserRepository(postgres.Connection)
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	setupTaskRoutes(taskRepo)
	setupUserRoutes(userRepo)

	http.ListenAndServe(":8080", nil)
}
