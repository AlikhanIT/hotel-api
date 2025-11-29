package main

import (
	_ "github.com/AlikhanIT/hotel-api/docs"
	"github.com/AlikhanIT/hotel-api/internal/adapter/db"
	"github.com/AlikhanIT/hotel-api/internal/adapter/http"
	"github.com/AlikhanIT/hotel-api/internal/adapter/server"
	"github.com/AlikhanIT/hotel-api/internal/config"
	domain "github.com/AlikhanIT/hotel-api/internal/domain/hotel"
	"github.com/AlikhanIT/hotel-api/internal/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	cfg := config.Load()

	logger.Info("application starting")

	dbInstance := db.NewPostgres(db.DBConfig{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
	})
	if err := dbInstance.AutoMigrate(&domain.Hotel{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	repo := db.NewHotelRepositoryGorm(dbInstance)
	handler := http.NewHandler(repo)
	router := server.NewRouter(handler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := server.Addr(cfg.Server.Port)
	logger.Info("server running on " + addr)
	_ = router.Run(addr)
}
