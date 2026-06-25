package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/PlinyTheYounger0/pliny_personal_portfolio/cmd/handlers"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "42069"
		logger.Warn("Port not set. Resorting to default: 42069")
	}

	filepathRoot := os.Getenv("FILEPATHROOT")
	if filepathRoot == "" {
		filepathRoot = "."
		logger.Warn("File Path Root not set. Resorting to default: .")
	}

	cfg := handlers.ApiConfig{
		Logger: logger,
	}

	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("GET /", cfg.Index)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Server Serving from %s & Listening on Port %s\n", filepathRoot, port)

	err := srv.ListenAndServe()
	if err != nil {
		logger.Error("Server Failed to Start", "error", err)
	}
}
