package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/joe-tripodi/chirpy/internal/database"
	"github.com/joe-tripodi/chirpy/internal/sanitize"
)

var ErrChirpTooLong = errors.New("Chirp is too long")

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserId    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string    `json:"body"`
		ID   uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}

	if params.ID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "user id must be present", nil)
	}

	cleaned, err := validateChirp(params.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   cleaned,
		UserID: params.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to save chirp", nil)
		return
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserId:    chirp.UserID,
	})

}

func validateChirp(body string) (string, error) {
	const maxChirpLength = 140

	if len(body) > maxChirpLength {
		return "", ErrChirpTooLong
	}
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned := sanitize.CleanChirp(body, badWords)
	return cleaned, nil
}
