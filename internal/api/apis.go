package api

import (
	"net/http"
	"task-process-service/internal/handler"
	db "task-process-service/internal/postgres"
	"task-process-service/internal/repository"
	"task-process-service/internal/service"

	httpSwagger "github.com/swaggo/http-swagger"
)

func StartServer() {
	taskService := service.NewTaskService(repository.NewTaskRepository(db.Connection))
	userService := service.NewUserService(repository.NewUserRepository(db.Connection))
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	distributeTasksHandler := handler.NewDistributeTasksHandler(taskService, userService)

	http.HandleFunc("/distributeTasks", distributeTasksHandler.DistributeTasks)
	setupTaskRoutes(taskService)
	setupUserRoutes(userService)

	http.ListenAndServe(":8080", nil)
}
