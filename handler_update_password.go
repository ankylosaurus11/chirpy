package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ankylosaurus11/chirpy/internal/auth"
	"github.com/ankylosaurus11/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUpdatePassword(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
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

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	updatedUser, err := cfg.db.UpdatePassword(r.Context(), database.UpdatePasswordParams{
		UpdatedAt:      time.Now(),
		Email:          params.Email,
		HashedPassword: hashedPassword,
		ID:             user,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user email/password", err)
		return
	}

	userResponse := User{
		ID:        user,
		UpdatedAt: updatedUser.UpdatedAt,
		Email:     updatedUser.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(userResponse)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error marshalling JSON", err)
		return
	}
	w.WriteHeader(200)
	w.Write(dat)
}
