package database

import (
	"os"

	"github.com/SandeshNilasKhatiwada/slack-clone/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func Connect() (*gorm.DB, error) {
	dsn := "host=" + env("DB_HOST", "localhost") +
		" user=" + env("DB_USER", "postgres") +
		" password=" + env("DB_PASSWORD", "sandesh1") +
		" dbname=" + env("DB_NAME", "pulsedb") +
		" port=" + env("DB_PORT", "5432") +
		" sslmode=disable"

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
