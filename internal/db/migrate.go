package db

import (
	"gorm.io/gorm"

	userDomain "qualthea-api/internal/app/user/domain"
)

// Migrate tables to the database
func MigrateTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		&userDomain.User{},
	)

	return err
}
