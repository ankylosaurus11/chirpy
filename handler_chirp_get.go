package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retreiving chirps", err)
		return
	}

	var fixedChirps []Chirp

	for _, chirp := range chirps {
		result := Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
		fixedChirps = append(fixedChirps, result)
	}

	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(fixedChirps)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error marshalling JSON", err)
		return
	}
	w.WriteHeader(200)
	w.Write(dat)
}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	chirpIDStr := r.PathValue("chirpID")

	chirpID, err := uuid.Parse(chirpIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	fixedChirp := Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}

	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(fixedChirp)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error marshalling JSON", err)
		return
	}
	w.WriteHeader(200)
	w.Write(dat)
}
