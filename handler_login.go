package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/joe-tripodi/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email            string        `json:"email"`
		Password         string        `json:"password"`
		ExpiresInSeconds time.Duration `json:"expires_in_seconds"`
	}

	var params parameters
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode params", err)
		return
	}

	const oneHour = time.Hour

	if params.ExpiresInSeconds == 0 || params.ExpiresInSeconds > oneHour {
		params.ExpiresInSeconds = oneHour
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "User does not exist", err)
		return
	}

	ok := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if ok != nil {
		respondWithError(w, http.StatusUnauthorized, "Password does not match", nil)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.secret, params.ExpiresInSeconds)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to create JWT", err)
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     token,
	})

}
