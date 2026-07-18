package api

import (
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"symbalx/api/endpoints"
	"symbalx/api/services"
)

var (
	backendKey   string
	spotifyAuth  string
	spotifyApi   string
	clientId     string
	clientSecret string
	databasePool *pgxpool.Pool
	jwtSecret    []byte
)

func loadEnvs() error {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		log.Fatal("Error loading .env file")
		return err
	}

	backendKey = os.Getenv("BACKEND_KEY")
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	spotifyAuth = os.Getenv("SPOTIFY_AUTH_API")
	spotifyApi = os.Getenv("SPOTIFY_API")
	clientId = os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	return nil
}

func InitializeServer(pool *pgxpool.Pool) error {
	err := loadEnvs()
	if err != nil {
		return err
	}

	databasePool = pool

	services.Init(databasePool, spotifyAuth, spotifyApi, clientId, clientSecret)
	endpoints.Init(databasePool, backendKey, jwtSecret)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
	return nil
}
