package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	logger := log.New(os.Stdout, "webapp ", log.LstdFlags)
	logger.Printf("Web app running on *:%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Request received for: %s\n", r.URL.Path)
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	server := &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Fatalf("Error starting the web server: %s\n", err)
	}
}
