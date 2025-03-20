package app

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"

	"qualthea-api/internal/db"
	"qualthea-api/internal/server"
)

// Run Starts the application
func Run() {
	// Load env variables
	err := godotenv.Load()

	if err != nil {
		slog.Error("Error loading .env file")
		os.Exit(1)
	}

	// Initialize database
	dbi, err := db.Init()

	// If there's an error
	if err != nil {
		slog.Error("Error initializing database: " + err.Error())
		os.Exit(1)
	}

	// Start server
	server.Start(dbi)
}
