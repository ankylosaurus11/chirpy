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

	mux.Handle("/", fs)

	fmt.Printf("Server starting on :%s...\n", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
