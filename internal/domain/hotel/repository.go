package domain

// Repository — чистый интерфейс. Не содержит имплементации.
type Repository interface {
	CreateHotel(h *Hotel) error
	GetHotels() ([]Hotel, error)
	GetHotelByID(id uint) (*Hotel, error)
	UpdateHotel(h *Hotel) error
	DeleteHotel(id uint) error
}
