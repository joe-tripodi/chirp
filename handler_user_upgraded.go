package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUserUpgraded(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserId uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	var params parameters
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to parse json", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = cfg.db.UpgradeUserToRed(r.Context(), params.Data.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "user not found", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "unable to update user", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
