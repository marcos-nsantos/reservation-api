package repository

import (
	"gorm.io/gorm"

	"github.com/marcos-nsantos/reservation-api/internal/entity"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) Create(reservation *entity.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *ReservationRepository) CheckAvailability(reservation *entity.Reservation) bool {
	var count int64
	r.db.Model(&entity.Reservation{}).
		Where("resource_id = ? AND ? < end_time AND ? > start_time AND status <> ?",
			reservation.ResourceID, reservation.EndTime, reservation.StartTime, entity.Cancelled).
		Count(&count)

	return count == 0
}
