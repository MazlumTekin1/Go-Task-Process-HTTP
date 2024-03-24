package main

import (
	"task-process-service/internal/api"
	postgres "task-process-service/internal/infrastructure"
)

func main() {
	postgres.Initialize()
	api.StartServer()
}
