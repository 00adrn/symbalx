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

func loadEnvs() (error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	connectionString = os.Getenv("DATABASE_URL")
	fmt.Println("Connection String: ", connectionString)
	return nil
}

func InitializeDatabase() (*pgxpool.Pool, error) {
	err := loadEnvs()
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return pool, nil
}


