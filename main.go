package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/current_time", currentTimeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
