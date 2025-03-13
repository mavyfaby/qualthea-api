package db

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

// Initialize database
func Init() (*gorm.DB, error) {
	var err error

	// Connect to database
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		),
	}))

  // If has error
  if err != nil {
    return nil, err
  }
	
  // Create the database if it doesn't exist
  err = db.Exec("CREATE DATABASE IF NOT EXISTS " + os.Getenv("DB_NAME")).Error

  // If error
  if err != nil {
    return nil, err
  }

  // Close the database connection
  err = Close()

  // If error
  if err != nil {
    return nil, err
  }

  // Connect to the database
  db, err = gorm.Open(mysql.New(mysql.Config{
    DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
      os.Getenv("DB_USER"),
      os.Getenv("DB_PASS"),
      os.Getenv("DB_HOST"),
      os.Getenv("DB_PORT"),
      os.Getenv("DB_NAME"),
    ),
  }))

  // If error
  if err != nil {
    return nil, err
  }

  // Set connection pool
  sqlDB, err := db.DB()

  // If error
  if err != nil {
    return nil, err
  }

  // Sets the maximum number of connections in the idle connection pool.
  sqlDB.SetMaxIdleConns(50)
  // Sets the maximum number of open connections to the database.
  sqlDB.SetMaxOpenConns(100)
  // Sets the maximum amount of time a connection may be reused.
  sqlDB.SetConnMaxLifetime(time.Hour)

  // Return the database connection
  return db, nil
}

// Close the database connection
func Close() error {
	// If the database connection is nil
	if db == nil {
		return errors.New("Can't close a nil database connection!")
	}

	// Get the underlying sql.DB object from the gorm.DB
	sqlDB, err := db.DB()

	// If error
	if err != nil {
		return err
	}

	// Close the database connection
	err = sqlDB.Close()

	// If error
	if err != nil {
		return err
	}

	return nil
}
