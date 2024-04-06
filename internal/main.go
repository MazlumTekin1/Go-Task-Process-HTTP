package main

import (
	"log"
	"task-process-service/internal/api"
	postgres "task-process-service/internal/postgres"
)

func main() {
	// Initialize the database connection
	postgres.Initialize()

	// Start the server
	log.Println("Starting the Go Task Process Service")
	api.StartServer()

}
