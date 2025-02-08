package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
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

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: params.Body,
	})
}
