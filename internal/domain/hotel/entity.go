package domain

import "time"

type Hotel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	Rating    float32   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
