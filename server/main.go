package main

import (
	"fmt"
	"symbalx/api"
	"symbalx/database"
)

func main() {

	pool, err := database.InitializeDatabase()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}
	defer pool.Close()

	api.InitializeServer(pool)
}
