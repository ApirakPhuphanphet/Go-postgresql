package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to pgAdmin database
func DatabaseConnection() *gorm.DB {
	dsn := "host=localhost user=posgres password=admin1234 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return db
	}
	return db
}
