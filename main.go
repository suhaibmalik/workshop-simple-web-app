package main

import (
	"log"
	"net/http"
)

// todo: add an endpoint for handling the user list
func main() {
	http.HandleFunc("/current_time", currentTimeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
