package config

import (
	"expense-tracker/routes"
	"fmt"
	"net/http"
	"time"
)

func StartServer() {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	server := http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	fmt.Println("Server starting on port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
