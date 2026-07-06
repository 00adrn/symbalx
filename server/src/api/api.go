package api

import (
	"context"
	"encoding/json"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"fmt"
	"io"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var (
	backendKey   string
	spotifyAuth  string
	spotifyApi   string
	clientId 	 string
	clientSecret string
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
	spotifyAuth = os.Getenv("SPOTIFY_AUTH_API")
	spotifyApi = os.Getenv("SPOTIFY_API")
	clientId = os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
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
	Username 	string `json:"username"`
	Email    	string `json:"email"`
	AccessToken string `json:"access_token"`
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

func readUserToken(r *http.Request) (string) {
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

func printHttpResponse(resp *http.Response) (error) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body")
		return err
	}

	bodyString := string(bodyBytes)
	log.Printf("Spotify API response: %s", bodyString)
	return nil
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
	err = databasePool.QueryRow(
		context.Background(), 
		"SELECT user_id, password_hash FROM users WHERE email = $1", 
		loginData.Email).Scan(&userID, &storedPassword)

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
	err = databasePool.QueryRow(
		context.Background(), 
		"INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING user_id", 
		data.Username, data.Email, string(hashedPassword)).Scan(&newUserID)

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
	err = databasePool.QueryRow(
		context.Background(), 
		"SELECT username, email FROM users WHERE user_id = $1", 
		userId).Scan(&userData.Username, &userData.Email)
	if err != nil {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}

	err = databasePool.QueryRow(
		context.Background(), 
		"SELECT spotify_token FROM spotify_tokens WHERE user_id = $1", 
		userId).Scan(&userData.AccessToken)
	if err != nil {
		log.Printf("Spotify Access token for %s not found", userData.Username)
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

	_, err = databasePool.Exec(
		context.Background(), 
		"INSERT INTO spotify_tokens (user_id, spotify_token, refresh_token) VALUES ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET spotify_token = EXCLUDED.spotify_token, refresh_token = EXCLUDED.refresh_token", 
		userId, data.AccessToken, data.RefreshToken)
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

func runTokenRefreshService() {

	for {
		log.Println("Starting refresh service...")

		rows, _ := databasePool.Query(
			context.Background(),
			"SELECT user_id, spotify_token, refresh_token FROM spotify_tokens",
		)

		var id, access, refresh string
		_, err := pgx.ForEachRow(rows, []any{&id, &access, &refresh}, func() error {
			body := url.Values{}
			body.Set("grant_type", "refresh_token")
			body.Set("refresh_token", refresh)
			body.Set("client_id", clientId)

			req, err := http.NewRequest("POST", spotifyAuth+"/token", strings.NewReader(body.Encode()))
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret)))

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println("Error making request token")
				return err
			}
			defer resp.Body.Close()

			var data SpotifyTokens
			err = json.NewDecoder(resp.Body).Decode(&data)
			if err != nil {
				log.Println("Error reading response body")
				return err
			}

			_, err = databasePool.Exec(
				context.Background(), 
				"UPDATE spotify_tokens SET spotify_token = $1 WHERE user_id = $2", 
				data.AccessToken, id)
			if err != nil {
				log.Println("Error updating the database")
				return err
			}

			if data.RefreshToken != "" {
				_, err = databasePool.Exec(
					context.Background(), 
					"UPDATE spotify_tokens SET refresh_token = $1 WHERE user_id = $2", 
					data.RefreshToken, id)
				if err != nil {
					log.Println("Error updating the database")
					return err
				}
			}

			log.Printf("Successfully refreshed token for user %s", id)
			return nil
		})
		if err != nil {
			log.Println("Error refreshing tokens")
			return
		}

		log.Println("Refresh service sleeping...")
		time.Sleep(30*60*time.Second)
	}
}

func runTrackHistoryService() {
	for {
		log.Println("Starting refresh service...")

		rows, _ := databasePool.Query(
			context.Background(),
			"SELECT user_id, spotify_token, last_checked FROM spotify_tokens",
		)

		var userId, spotifyToken string
		var lastChecked time.Time
		_, err := pgx.ForEachRow(rows, []any{&userId, &spotifyToken, &lastChecked}, func() error {

			req, err := http.NewRequest("GET", spotifyApi+"/me/player/recently-played?after=" + fmt.Sprintf("%d", lastChecked.UnixMilli()), nil)
			if err != nil {
				log.Println("Error creating request")
				return err
			}

			req.Header.Set("Authorization", "Bearer " + spotifyToken)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println("Error making request to Spotify API")
				return err
			}
			defer resp.Body.Close()

			err = printHttpResponse(resp)
			if err != nil {
				log.Println("Error printing response")
				return err
			}
			
			if resp.StatusCode != http.StatusOK {
				log.Printf("Spotify API returned status code %d", resp.StatusCode)
				return nil
			}

			return nil
		})
		if err != nil {
			log.Println("Error running track history service")
			log.Println(err)
			return
		}

		log.Println("Track history service sleeping...")
		time.Sleep(30*60*time.Second)
	}
}

func InitializeServer(pool *pgxpool.Pool) error {
	err := loadEnvs()
	if err != nil {
		return err
	}

	databasePool = pool
	go runTokenRefreshService()
	go runTrackHistoryService()
	http.HandleFunc("/auth/login", accountLogin)
	http.HandleFunc("/auth/register", accountRegister)
	http.HandleFunc("/user/profile", getProfile)
	http.HandleFunc("/user/spotify/update", updateSpotifyInfo)
	http.ListenAndServe(":8080", nil)
	return nil
}
