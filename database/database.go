package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// return a connection pool
func ConnectDB() *pgxpool.Pool {
	url := os.Getenv("DB_URL")
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(),config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	fmt.Println("Connected to database")
	return dbPool
}