package main

import (
	"task-process-service/internal/api"
	postgres "task-process-service/internal/postgres"
)

func main() {
	// Initialize the database connection
	postgres.Initialize()

	// Start the server
	api.StartServer()

}
