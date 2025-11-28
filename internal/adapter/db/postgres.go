package db

import (
	"fmt"

	domain "github.com/AlikhanIT/hotel-api/internal/domain/hotel"
	"github.com/AlikhanIT/hotel-api/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

func NewPostgres(cfg DBConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("cannot connect to postgres", err)
		panic(err)
	}

	if err := db.AutoMigrate(&domain.Hotel{}); err != nil {
		logger.Error("failed to migrate auto", err)
		panic(err)
	}

	logger.Info("connected to postgres")
	return db
}
