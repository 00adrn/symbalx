package main;

import (
	"net/http"
	"os"
	"fmt"
);



func main() {
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil);
}