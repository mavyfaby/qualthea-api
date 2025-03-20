package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Declare a global variable to holds the database connection
var db *sql.DB

// Init database connection
func Init() (*sql.DB, error) {
	// Load environment variables
	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")

	// Declare an error variable
	var err error

	// Connect to database
	db, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4",
			DbUser, DbPass, DbHost, DbPort,
		),
	)

	// If there is an error connecting to the database, return the error
	if err != nil {
		return nil, err
	}

	// Create the database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + DbName)

	// If there is an error creating the database, return the error
	if err != nil {
		return nil, err
	}

	// Close the connection to the database
	err = db.Close()

	// If there is an error closing the connection to the database, return the error
	if err != nil {
		return nil, err
	}

	// Connect to the database
	db, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
			DbUser, DbPass, DbHost, DbPort, DbName,
		),
	)

	// If there is an error connecting to the database, return the error
	if err != nil {
		return nil, err
	}

	// Sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(50)
	// Sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(100)
	// Sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(time.Hour)

	// Return the database connection
	return db, nil
}

// Close a database connection
func Close() error {
	// If the database connection is nil
	if db == nil {
		return errors.New("can't close a nil database connection")
	}

	// Close the connection to the database
	return db.Close()
}
