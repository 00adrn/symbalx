package database

import (
	"fmt"
	"log"
	"os"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	connectionString string
	databasePool *pgxpool.Pool
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString = os.Getenv("DATABASE_URL")
	fmt.Println("Connection String: ", connectionString)
}

func InitClient() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}


