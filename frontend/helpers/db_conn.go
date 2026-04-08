package helpers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConn() *gorm.DB {
	if DB != nil {
		return DB
	}

	dsn := "host=localhost user=ide_user password=password123 dbname=ide_database port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to DB: %v", err))
	}

	DB = db
	return DB
}
