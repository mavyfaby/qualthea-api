package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"qualthea-api/internal/db/models/category"
	"qualthea-api/internal/db/models/user"

	"qualthea-api/internal/server/handlers"

	categoryDb "qualthea-api/internal/db/models/category/db"
	userDb "qualthea-api/internal/db/models/user/db"
)

// Start the app server
func Start(db *sql.DB) {
	// Create a new echo server
	server := echo.New()

	// Hide the server banner
	server.HideBanner = true
	server.HidePort = true

	// Setup middleware for the server
	server.Pre(middleware.RemoveTrailingSlash())

	// Get port from environment variable
	DbPort := os.Getenv("PORT")

	// Create services
	categoryService := category.NewService(categoryDb.New(db))
	userService := user.NewService(userDb.New(db))

	// Bind handlers
	handlers.CategoryHandler(server, categoryService)
	handlers.UserHandler(server, userService)

	// Run the server in a goroutine
	go func() {
		// Start the server on port 3000
		// If an error occurs except for the server being closed
		if err := server.Start(":" + DbPort); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// Log that the server failed to start
			fmt.Println("🟥 " + err.Error())
			// Quit the application
			os.Exit(1)
		}
	}()

	// Print that the server has started
	fmt.Println("✅ Qualthea Bookstore API Server started on port " + DbPort)

	// Create a channel to listen for an OS signal
	sigterm := make(chan os.Signal, 1)
	// Notify the channel when an Interrupt signal is received
	signal.Notify(sigterm, os.Interrupt)
	// Block the main thread until an Interrupt signal is received
	<-sigterm

	// Print that the server is shutting down
	fmt.Println("Gracefully shutting down...")

	// Close db connection
	if err := db.Close(); err != nil {
		slog.Error("Failed to close database connection: " + err.Error())
	}

	// Gracefully shutdown the server
	if err := server.Shutdown(context.Background()); err != nil {
		slog.Error("Failed to shutdown server gracefully: " + err.Error())
	}

	// Print that the server has been stopped
	fmt.Println("🛑 Qualthea Bookstore API Server stopped.")
}
