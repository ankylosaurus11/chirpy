package main

import (
	"fmt"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("."))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/app/", http.StripPrefix("/app", fs))
	mux.HandleFunc("/healthz", handlerReadiness)

	fmt.Printf("Server starting on :%s...\n", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
