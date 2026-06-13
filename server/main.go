package main

import (
	"context"
	"fmt"
	"symbalx/api"
	"symbalx/database"
)

func main() {
	fmt.Println("Hello, World!")
	api.TestFunc()
	database.LoadEnvs()

	pool, err := database.InitClient()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}
	defer pool.Close()

	var username string
	var email string

	err = pool.QueryRow(context.Background(), "SELECT username, email FROM users WHERE username = $1 AND email = $2", "adrian", "adrianvez11@gmail.com").Scan(&username, &email)
	if err != nil {
		fmt.Println("Error querying database: ", err)
		return
	}
	fmt.Printf("Read Values: %s, %s\n", username, email)
}
