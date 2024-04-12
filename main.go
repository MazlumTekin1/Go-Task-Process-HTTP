package main

import (
	"fmt"
	"log"
	"path/filepath"
	"task-process-service/internal/api"
	postgres "task-process-service/internal/postgres"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Task Process Service API
// @version 1.0
// @description This is a Task Process Service.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:45009
// @BasePath /internal/api
// @schemes http
func main() {
	// absPath, err := filepath.Abs("../internal/config.json")
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
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:45009/swagger/doc.json"),
	))
	api.StartServer(r)

}
