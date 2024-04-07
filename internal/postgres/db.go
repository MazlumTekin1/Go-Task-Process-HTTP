package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Connection *pgxpool.Pool

func Initialize(absPath string) {

	cfg, err := LoadConfig(absPath)
	if err != nil {
		log.Fatalf("Error loading the configuration file: %v", err)
	}

	// PostgreSQL Connection
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s pool_max_conns=%d",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode,
		cfg.Database.MaxPoolSize)
	Connection, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Failed to create PostgreSQL connection: %v", err)
		os.Exit(1)
	}
	// Bağlantıyı test et
	if err := Connection.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to test PostgreSQL connection: %v", err)
		os.Exit(1)
	} else {
		log.Println("Database Connection Created")
	}
}
