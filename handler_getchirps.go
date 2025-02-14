package main

import (
	"encoding/json"
	"net/http"
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
