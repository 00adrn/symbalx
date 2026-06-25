package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
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

func loadEnvs() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	backendKey = os.Getenv("BACKEND_KEY")
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	return nil
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

type ProfileData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	// ProfilePicture string `json:"profile_picture"`
}

type SpotifyTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func prepareResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
}

func checkAuthorization(w http.ResponseWriter, r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return false
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
		return false
	}

	if token != backendKey {
		http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
		return false
	}

	return true
}

func generateToken(userId string, expirationTime time.Time) (string, error) {

	claims := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func readIdFromToken(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	return claims.UserID, nil
}

func readUserToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		return ""
	}

	return token
}

func accountLogin(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s on /auth/login", r.Method)
	if r.Method != http.MethodPost {
		prepareResponse(w, http.StatusMethodNotAllowed)
		return
	}

	if !checkAuthorization(w, r) {
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
	tokenString, err := generateToken(userID, expirationTime)
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
	log.Printf("%s on /auth/register", r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if !checkAuthorization(w, r) {
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
	tokenString, err := generateToken(newUserID, expirationTime)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Token signage error:", err)
		return
	}

	prepareResponse(w, http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString, "expires": expirationTime.Format(time.RFC3339)})
	log.Println("User successfully registered")
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s on /user/profile", r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userToken := readUserToken(r)
	userId, err := readIdFromToken(userToken)
	if err != nil {
		http.Error(w, "Invalid JWT Token", http.StatusNotAcceptable)
		return
	}

	var userData ProfileData
	err = databasePool.QueryRow(context.Background(), "SELECT username, email FROM users WHERE user_id = $1", userId).Scan(&userData.Username, &userData.Email)
	if err != nil {
		log.Println("Error: Non-real user token")
	}
	log.Printf("Read user data for %s\n", userData.Username)

	prepareResponse(w, http.StatusOK)
	err = json.NewEncoder(w).Encode(userData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func updateSpotifyInfo(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s on /spotify/update", r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userToken := readUserToken(r)
	userId, err := readIdFromToken(userToken)
	if err != nil {
		http.Error(w, "Error: Missing or invalid user token", http.StatusNotAcceptable)
		return
	}

	var data SpotifyTokens
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("Error reading body")
		return
	}

	_, err = databasePool.Exec(context.Background(), "INSERT INTO spotify_tokens (user_id, spotify_token, refresh_token) VALUES ($1, $2, $3)", userId, data.AccessToken, data.RefreshToken)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error updating spotify tokens:", err)
		return
	}

	prepareResponse(w, http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	if err != nil {
		log.Println("Error encoding response:", err)
	}
}

func InitializeServer(pool *pgxpool.Pool) error {
	err := loadEnvs()
	if err != nil {
		return err
	}

	databasePool = pool
	http.HandleFunc("/auth/login", accountLogin)
	http.HandleFunc("/auth/register", accountRegister)
	http.HandleFunc("/user/profile", getProfile)
	http.HandleFunc("/user/spotify/update", updateSpotifyInfo)
	http.ListenAndServe(":8080", nil)
	return nil
}
