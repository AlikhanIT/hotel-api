package db

import (
	domain "github.com/AlikhanIT/hotel-api/internal/domain/hotel"
	"github.com/AlikhanIT/hotel-api/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
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
