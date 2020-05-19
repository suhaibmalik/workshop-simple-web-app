package main

import (
	"log"
	"net/http"
)

// todo: create a file for handling the user list
func main() {
	http.HandleFunc("/current_time", currentTimeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
