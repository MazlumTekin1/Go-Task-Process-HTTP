package api

import (
	"net/http"
	postgres "task-process-service/internal/infrastructure"
	"task-process-service/internal/repository"
)

func StartServer() {
	taskRepo := repository.NewTaskRepository(postgres.Connection)
	userRepo := repository.NewUserRepository(postgres.Connection)

	setupTaskRoutes(taskRepo)
	setupUserRoutes(userRepo)

	http.ListenAndServe(":8080", nil)
}
