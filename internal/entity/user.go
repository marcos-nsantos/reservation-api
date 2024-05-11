package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint64         `json:"id"`
	Name         string         `gorm:"type:varchar(255);not null" json:"name"`
	Email        string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password     string         `gorm:"type:varchar(255);not null" json:"-"`
	Reservations []Reservation  `json:"reservations"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
