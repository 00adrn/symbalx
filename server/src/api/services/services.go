package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	databasePool *pgxpool.Pool
	spotifyAuth  string
	spotifyApi   string
	clientId     string
	clientSecret string
	readLock     bool
)

type SpotifyTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type spotifyTrackUri struct {
	Uri string `json:"uri"`
}

type spotifyTrackItem struct {
	Track spotifyTrackUri `json:"track"`
}

type SpotifyTrackHistory struct {
	Items []spotifyTrackItem `json:"items"`
}

func Init(pool *pgxpool.Pool, auth string, api string, cid string, csecret string) {
	databasePool = pool
	spotifyAuth = auth
	spotifyApi = api
	clientId = cid
	clientSecret = csecret

	go runTokenRefreshService()
	go runTrackHistoryService()
}

func runTokenRefreshService() {
	for {
		if readLock {
			log.Println("Refresh service is locked, waiting...")
			time.Sleep(10 * time.Second)
			continue
		}
		readLock = true
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

			req, err := http.NewRequest("POST", spotifyAuth+"/token", strings.NewReader(body.Encode()))
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret)))

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
			readLock = false
			return
		}

		log.Println("Refresh service sleeping...")
		readLock = false
		time.Sleep(30 * 60 * time.Second)
	}
}

func runTrackHistoryService() {
	for {
		if readLock {
			log.Println("Track history service is locked, waiting...")
			time.Sleep(10 * time.Second)
			continue
		}
		readLock = true
		log.Println("Starting refresh service...")

		rows, _ := databasePool.Query(
			context.Background(),
			"SELECT user_id, spotify_token, last_checked FROM spotify_tokens",
		)

		_, err := databasePool.Exec(
			context.Background(),
			"UPDATE spotify_tokens SET last_checked = NOW()",
		)
		if err != nil {
			log.Println("Error updating last_checked")
			readLock = false
			return
		}

		var userId, spotifyToken string
		var lastChecked time.Time
		_, err = pgx.ForEachRow(rows, []any{&userId, &spotifyToken, &lastChecked}, func() error {

			req, err := http.NewRequest("GET", spotifyApi+"/me/player/recently-played?after="+fmt.Sprintf("%d", lastChecked.UnixMilli()), nil)
			if err != nil {
				log.Println("Error creating request")
				return err
			}

			req.Header.Set("Authorization", "Bearer "+spotifyToken)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println("Error making request to Spotify API")
				return err
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				log.Printf("Spotify API returned status code %d", resp.StatusCode)
				return fmt.Errorf("Spotify API returned status code %d", resp.StatusCode)
			}

			var trackHistory SpotifyTrackHistory
			err = json.NewDecoder(resp.Body).Decode(&trackHistory)
			if err != nil {
				log.Println("Error decoding track history")
				return err
			}

			for i, item := range trackHistory.Items {
				log.Printf("Track %d: %s", i, item.Track.Uri)
				_, err = databasePool.Exec(
					context.Background(),
					"INSERT INTO spotify_track_history (user_id, track_id, date_listened) VALUES ($1, $2, NOW())",
					userId, item.Track.Uri,
				)
				if err != nil {
					log.Println("Error inserting track into database")
					return err
				}
			}

			return nil
		})
		if err != nil {
			log.Println("Error running track history service")
			readLock = false
			log.Println(err)
			return
		}

		log.Println("Track history service sleeping...")
		readLock = false
		time.Sleep(10 * 60 * time.Second)
	}
}
