package database

import (
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func connectDB() *pgxpool.Pool {
	url := os.Getenv("DB_URL")

	return db
}