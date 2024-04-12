package main

import (
	"fmt"
	"log"
	"path/filepath"
	"task-process-service/docs"
	"task-process-service/internal/api"
	postgres "task-process-service/internal/postgres"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for Swagger documentation
// @host localhost:45009
// @BasePath /
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
	r := chi.NewRouter()
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:45009"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:45009/swagger/doc.json"),
	))

	api.StartServer(r)
}
