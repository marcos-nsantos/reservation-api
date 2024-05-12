package entity

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	Pending   Status = "pending"
	Approved  Status = "approved"
	Rejected  Status = "rejected"
	Cancelled Status = "cancelled"
)

type Reservation struct {
	ID         uint64         `json:"id"`
	UserID     uint64         `json:"userID"`
	ResourceID uint64         `json:"resourceID"`
	Resource   Resource       `json:"-"`
	StartTime  time.Time      `gorm:"not null" json:"startTime"`
	EndTime    time.Time      `gorm:"not null" json:"endTime"`
	Status     Status         `gorm:"type:varchar(9);not null;default:'pending" json:"status"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
