package main

import (
	"fmt"
	_ "github.com/AlikhanIT/hotel-api/docs"
	"github.com/AlikhanIT/hotel-api/internal/adapter/db"
	"github.com/AlikhanIT/hotel-api/internal/adapter/http"
	"github.com/AlikhanIT/hotel-api/internal/adapter/server"
	"github.com/AlikhanIT/hotel-api/internal/config"
	"github.com/AlikhanIT/hotel-api/internal/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.Load()

	logger.Info("application starting")

	dbInstance := db.NewPostgres(cfg.DB.Url)
	repo := db.NewHotelRepositoryGorm(dbInstance)
	handler := http.NewHandler(repo)
	router := server.NewRouter(handler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Info("server running on " + addr)
	_ = router.Run(addr)
}
