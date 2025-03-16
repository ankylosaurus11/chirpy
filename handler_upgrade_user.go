package main

import (
	"encoding/json"
	"net/http"

	"github.com/ankylosaurus11/chirpy/internal/auth"
	"github.com/ankylosaurus11/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUpgradeUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error retreiving API Key", err)
		return
	}
	if apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized API Key", err)
		return
	}

	err = cfg.db.UpgradeRed(r.Context(), database.UpgradeRedParams{
		IsChirpyRed: true,
		ID:          params.Data.UserID,
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
