package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint64
	Name         string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"type:varchar(255);unique;not null"`
	Password     string `gorm:"type:varchar(255);not null"`
	Reservations []Reservation
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
