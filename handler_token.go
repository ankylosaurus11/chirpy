package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ankylosaurus11/chirpy/internal/auth"
	"github.com/ankylosaurus11/chirpy/internal/database"
)

func (cfg *apiConfig) handlerRefreshToken(w http.ResponseWriter, r *http.Request) {
	type Token struct {
		Token string `json:"token"`
	}
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retreive Token", err)
		return
	}

	refreshToken, err := cfg.db.GetToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Token does not exist", err)
		return
	}

	if refreshToken.ExpiresAt.Before(time.Now()) || refreshToken.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Token has been revoked or is expired", err)
		return
	}

	accessToken, err := auth.MakeJWT(refreshToken.UserID, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create JWT", err)
		return
	}

	newToken := Token{
		Token: accessToken,
	}

	respondWithJSON(w, 200, newToken)
}

func (cfg *apiConfig) handlerRevokeToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retreive Token", err)
		return
	}

	refreshToken, err := cfg.db.GetToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Token does not exist", err)
		return
	}

	cfg.db.UpdateToken(r.Context(), database.UpdateTokenParams{
		UpdatedAt: time.Now(),
		RevokedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Token: refreshToken.Token,
	})

	w.WriteHeader(http.StatusNoContent)
}
