package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func validateChirp(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Body string `json:"body"`
	}
	type validResponse struct {
		Valid bool `json:"valid"`
	}
	type errorResponse struct {
		Error string `json:"error"`
	}

	errorDefault := errorResponse{
		Error: "Something went wrong",
	}

	errorLength := errorResponse{
		Error: "Chirp is too long",
	}

	valid := validResponse{
		Valid: true,
	}

	decoder := json.NewDecoder(r.Body)
	body := requestBody{}
	err := decoder.Decode(&body)

	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		dat, err := json.Marshal(errorDefault)
		if err != nil {
			log.Printf("Error marshalling error response: %s", err)
			return
		}
		w.Write(dat)
		return
	}
	if len(body.Body) > 140 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		dat, err := json.Marshal(errorLength)
		if err != nil {
			log.Printf("Error marshalling error response: %s", err)
			return
		}
		w.Write(dat)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	dat, err := json.Marshal(valid)
	if err != nil {
		log.Printf("Error marshalling error response: %s", err)
		return
	}
	w.Write(dat)
}
