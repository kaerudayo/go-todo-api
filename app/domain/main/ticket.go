package domain

import "time"

type Ticket struct {
	ID          int       `gorm:"primaryKey"`
	EpicId      int       `gorm:"column:epic_id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
