package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	auth "qualthea-api/internal/app/auth/api"
	database "qualthea-api/internal/db"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Main application entry point
func Run() {
	// Load env variables
	err := godotenv.Load()

	if err != nil {
		slog.Error("Error loading .env file")
		os.Exit(1)
	}

	// Initialize database
	db, err := database.Init()

	// if error
	if err != nil {
		slog.Error("Error initializing database: " + err.Error())
		os.Exit(1)
	}

	// Migrate tables
	err = database.MigrateTables(db)

	// if error
	if err != nil {
		slog.Error("Error migrating tables: " + err.Error())
		os.Exit(1)
	}

	// Create a new Echo server
	server := echo.New()
	// Hide the server banner
	server.HideBanner = true
	server.HidePort = true

	// Setup middleware for the server
	server.Pre(middleware.RemoveTrailingSlash())

	// Setup the routes for the application
	auth.RegisterRoutes(server)

	// Run the server in a goroutine
	go func() {
		// Start the server on port 3000
		// If an error occurs except for the server being closed
		if err := server.Start(":3000"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// Log that the server failed to start
			fmt.Println("🟥 " + err.Error())
			// Quit the application
			os.Exit(1)
		}
	}()

	// Print that the server has started
	fmt.Println("✅ Qualthea Bookstore API Server started on port 3000")

	// Create a channel to listen for an OS signal
	sigterm := make(chan os.Signal, 1)
	// Notify the channel when an Interrupt signal is received
	signal.Notify(sigterm, os.Interrupt)
	// Block the main thread until an Interrupt signal is received
	<-sigterm

	// Print that the server is shutting down
	fmt.Println("Gracefully shutting down...")

	// Close db connection
	if err := database.Close(); err != nil {
		slog.Error("Failed to close database connection: " + err.Error())
	}

	// Gracefully shutdown the server
	if err := server.Shutdown(context.Background()); err != nil {
		slog.Error("Failed to shutdown server gracefully: " + err.Error())
	}

	// Print that the server has been stopped
	fmt.Println("🛑 Qualthea Bookstore API Server stopped.")
}
