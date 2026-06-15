package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var (
	backendKey   string
	databasePool *pgxpool.Pool
	jwtSecret    []byte
)

func loadEnvs() (error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	backendKey = os.Getenv("BACKEND_KEY")
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	return nil
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

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type RegisterData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func accountLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		prepareResponse(w, http.StatusMethodNotAllowed)
		return
	}

	var loginData LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var storedPassword string
	var userID string
	err = databasePool.QueryRow(context.Background(), "SELECT user_id, password_hash FROM users WHERE email = $1", loginData.Email).Scan(&userID, &storedPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error fetching user:", err)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(loginData.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error signing token:", err)
		return
	}

	prepareResponse(w, http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	log.Println("User logged in successfully")
}

func accountRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	authorization := r.Header.Get("Authorization")
	if authorization != backendKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var data RegisterData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Println("Attempting to create user with email and username: ", data.Email, data.Username)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error hashing password:", err)
		return
	}

	var newUserID string
	err = databasePool.QueryRow(context.Background(), "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING user_id", data.Username, data.Email, string(hashedPassword)).Scan(&newUserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			http.Error(w, "Username or email already exists", http.StatusConflict)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("User insertion error:", err)
		}
		return
	}

	log.Println("User created successfully with ID:", newUserID)

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: newUserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Token signage error:", err)
		return
	}

	prepareResponse(w, http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString, "expires": expirationTime.Format(time.RFC3339)})
	log.Println("User successfully registered")
}

func InitializeServer(pool *pgxpool.Pool) (error) {
	err := loadEnvs()
	if err != nil {
		return err
	}

	databasePool = pool
	http.HandleFunc("/auth/login", accountLogin)
	http.HandleFunc("/auth/register", accountRegister)
	http.ListenAndServe(":8080", nil)
	return nil
}
