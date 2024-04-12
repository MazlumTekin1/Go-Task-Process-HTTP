package main

import (
	"fmt"
	"log"
	"path/filepath"
	"task-process-service/internal/api"
	postgres "task-process-service/internal/postgres"
)

func main() {
	absPath, err := filepath.Abs("./internal/config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(absPath)

	// Initialize the database connection
	postgres.Initialize(absPath)

	// Start the server
	log.Println("Starting the Go Task Process Service!")

	api.StartServer()
}
