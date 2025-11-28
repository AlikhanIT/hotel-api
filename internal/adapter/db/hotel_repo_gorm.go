package db

import (
	domain "github.com/AlikhanIT/hotel-api/internal/domain/hotel"
	"github.com/AlikhanIT/hotel-api/internal/logger"
	"gorm.io/gorm"
)

type HotelRepositoryGorm struct {
	db *gorm.DB
}

// Конструктор
func NewHotelRepositoryGorm(db *gorm.DB) domain.Repository {
	return &HotelRepositoryGorm{db: db}
}

func (r *HotelRepositoryGorm) CreateHotel(h *domain.Hotel) error {
	if err := r.db.Create(h).Error; err != nil {
		logger.Error("failed to create hotel", err)
		return err
	}
	logger.Info("hotel created")
	return nil
}

func (r *HotelRepositoryGorm) GetHotels() ([]domain.Hotel, error) {
	var hotels []domain.Hotel
	if err := r.db.Find(&hotels).Error; err != nil {
		logger.Error("failed to get all hotels", err)
		return nil, err
	}
	return hotels, nil
}

func (r *HotelRepositoryGorm) GetHotelByID(id uint) (*domain.Hotel, error) {
	var h domain.Hotel
	if err := r.db.First(&h, id).Error; err != nil {
		logger.Error("hotel not found", err)
		return nil, err
	}
	return &h, nil
}

func (r *HotelRepositoryGorm) UpdateHotel(h *domain.Hotel) error {
	if err := r.db.Save(h).Error; err != nil {
		logger.Error("failed to update hotel", err)
		return err
	}
	logger.Info("hotel updated")
	return nil
}

func (r *HotelRepositoryGorm) DeleteHotel(id uint) error {
	if err := r.db.Delete(&domain.Hotel{}, id).Error; err != nil {
		logger.Error("failed to delete hotel", err)
		return err
	}
	logger.Info("hotel deleted")
	return nil
}
