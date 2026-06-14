package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	backendKey   string
	databasePool *pgxpool.Pool
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	backendKey = os.Getenv("BACKEND_KEY")
}

func prepareResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func accountLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		prepareResponse(w, http.StatusMethodNotAllowed)
		return
	}

	authorization := r.Header.Get("Authorization")
	if authorization != backendKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var loginData LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	prepareResponse(w, http.StatusOK)

	err = json.NewEncoder(w).Encode(loginData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("read %v and %v", loginData.Email, loginData.Password)

}

type RegisterData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func accountRegister(w http.ResponseWriter, r *http.Request) {
}

func InitializeServer(pool *pgxpool.Pool) {
	databasePool = pool
	http.HandleFunc("/auth/login", accountLogin)
	http.HandleFunc("/auth/register", accountRegister)
	http.ListenAndServe(":8080", nil)
}
