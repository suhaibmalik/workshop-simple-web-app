package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// todo: replace CurrentTime with the int type
type CurrentTimeResponse struct {
	CurrentTime string
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	ctr := CurrentTimeResponse{
		CurrentTime: "0",
	}

	js, err := json.Marshal(ctr)
	if err != nil {
		log.Fatal("couldn't parse to json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
}
