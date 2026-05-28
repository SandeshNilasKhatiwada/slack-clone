package database

import (
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=sandesh1 dbname=pulsedb port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// auto migrate the database schema
	conn.AutoMigrate(&models.User{})
	db = conn
	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
