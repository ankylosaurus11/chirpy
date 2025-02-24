package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ankylosaurus11/chirpy/internal/auth"
	"github.com/ankylosaurus11/chirpy/internal/database"
)

func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Authorization header is missing or malformed", err)
		return
	}

	user, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid or expired token", err)
		return
	}

	if params.Body == "" {
		respondWithError(w, http.StatusBadRequest, "Chirp cannot be empty", nil)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(params.Body, " ")
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		for _, badWord := range badWords {
			if lowerWord == badWord {
				words[i] = "****"
			}
		}
	}
	params.Body = strings.Join(words, " ")

	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   params.Body,
		UserID: user,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating chirp", err)
		return
	}
	chirpResponse := Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(chirpResponse)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error marshalling JSON", err)
		return
	}
	w.WriteHeader(201)
	w.Write(dat)
}
