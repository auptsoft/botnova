// internal/storage/db.go
package gorm

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Driver string // "postgres" or "sqlite"
	DSN    string
}

func NewDB(cfg Config) *gorm.DB {
	var db *gorm.DB
	var err error

	switch cfg.Driver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
	default:
		log.Fatalf("unsupported driver: %s", cfg.Driver)
	}

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}
