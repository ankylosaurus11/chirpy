package main

import (
	"fmt"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Server starting on :%s...\n", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
