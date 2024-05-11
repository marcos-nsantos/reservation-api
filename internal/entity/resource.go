package entity

import "time"

type Resource struct {
	ID          uint64
	Name        string `gorm:"type:varchar(255);not null"`
	Capacity    uint32 `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
