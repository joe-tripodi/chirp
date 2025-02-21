package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/joe-tripodi/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {

	chirpId, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to parse chirpId", err)
		return
	}
	chirp, err := cfg.db.GetChirp(r.Context(), chirpId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp does not exist", err)
		return
	}
	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "no access token", err)
		return
	}
	userId, err := auth.ValidateJWT(accessToken, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "malformed access token", err)
		return
	}

	if chirp.UserID != userId {
		respondWithError(w, http.StatusForbidden, "you do not have access to delete this chirp", err)
		return
	}

	err = cfg.db.DeleteChirpById(r.Context(), chirpId)

	if err != nil {
		respondWithError(w, http.StatusForbidden, "you do not have access to delete this chirp", err)
		return
	}

	w.WriteHeader(204)
	return
}
