package db

import (
	"karangiri121/railway-scrapper/internal/config"
	"log/slog"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func New(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.StoragePath), &gorm.Config{})

	if err != nil {
		slog.Info("Error in Opening DB", slog.String("Error", err.Error()))
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	_, err = sqlDB.Exec(`PRAGMA journal_mode=WAL;`)
	if err != nil {
		return nil, err
	}

	return db, err
}
