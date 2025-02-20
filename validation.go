package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("failed to parse json: %s", err)
		errMsg := fmt.Sprintf(`{"error": "%s"}`, err)
		w.WriteHeader(500)
		w.Write([]byte(errMsg))
		return
	}

	if len(params.Body) > 140 {
		log.Println("Invalid tweet length, larger than 140")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte(`{"valid": false}`))
		return
	}

	log.Println("Valid tweet")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"valid": true}`))
}
