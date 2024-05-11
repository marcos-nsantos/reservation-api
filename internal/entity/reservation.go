package entity

import "time"

type Status string

const (
	Pending   Status = "pending"
	Approved  Status = "approved"
	Rejected  Status = "rejected"
	Cancelled Status = "cancelled"
)

type Reservation struct {
	ID         uint64
	UserID     uint64
	ResourceID uint64
	Resource   Resource
	StartTime  time.Time `gorm:"not null"`
	EndTime    time.Time `gorm:"not null"`
	Status     Status    `gorm:"type:varchar(9);not null;default:'pending"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
