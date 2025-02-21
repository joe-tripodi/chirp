package main

import (
	"net/http"

	"github.com/joe-tripodi/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	headerRefreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "refresh token is not present", err)
		return
	}

	err = cfg.db.RevokeRefreshToken(r.Context(), headerRefreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to revoke refresh token", err)
		return
	}

	w.WriteHeader(204)

}
