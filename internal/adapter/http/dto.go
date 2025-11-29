package http

import domain "github.com/AlikhanIT/hotel-api/internal/domain/hotel"

type HotelDTO struct {
	Name    string  `json:"name" binding:"required"`
	City    string  `json:"city" binding:"required"`
	Address string  `json:"address" binding:"required"`
	Rating  float32 `json:"rating" binding:"required"`
}

// Конвертация DTO → доменная модель
func (dto HotelDTO) ToDomain() *domain.Hotel {
	return &domain.Hotel{
		Name:    dto.Name,
		City:    dto.City,
		Address: dto.Address,
		Rating:  dto.Rating,
	}
}
