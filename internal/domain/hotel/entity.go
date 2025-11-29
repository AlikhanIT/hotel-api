package domain

import "time"

type Hotel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	Rating    float32   `json:"rating"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
