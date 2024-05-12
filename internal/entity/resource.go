package entity

import (
	"time"

	"gorm.io/gorm"
)

type Resource struct {
	ID          uint64         `json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Capacity    uint32         `gorm:"not null" json:"capacity"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
