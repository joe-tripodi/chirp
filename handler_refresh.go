package main

import (
	"net/http"
	"time"

	"github.com/joe-tripodi/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}
	headerRefreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "refresh token is not present", err)
		return
	}

	refreshToken, err := cfg.db.GetRefreshToken(r.Context(), headerRefreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "refresh token is expired", err)
	}

	jwtToken, err := auth.MakeJWT(refreshToken.UserID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to generate jwt token", err)
	}

	respondWithJSON(w, http.StatusOK, response{
		Token: jwtToken,
	})

}
