package main

import (
	"github.com/AlikhanIT/hotel-api/internal/adapter/db"
	"github.com/AlikhanIT/hotel-api/internal/adapter/http"
	"github.com/AlikhanIT/hotel-api/internal/adapter/server"
	"github.com/AlikhanIT/hotel-api/internal/config"
	"github.com/AlikhanIT/hotel-api/internal/logger"
)

func main() {
	cfg := config.Load()

	logger.InitSeq(logger.SeqConfig{
		URL:     cfg.Logging.SeqURL,
		APIKey:  cfg.Logging.SeqAPIKey,
		Enabled: cfg.Logging.Enabled,
	})
	logger.Info("application starting")

	dbInstance := db.NewPostgres(db.DBConfig{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
	})
	repo := db.NewHotelRepositoryGorm(dbInstance)
	handler := http.NewHandler(repo)
	router := server.NewRouter(handler)

	addr := server.Addr(cfg.Server.Port)
	logger.Info("server running on " + addr)
	_ = router.Run(addr)
}
